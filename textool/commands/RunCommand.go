package commands

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/pkg/pathutil"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/go-akka/configuration"
	"github.com/gookit/color"
	"github.com/moby/term"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}
type DockerMessage struct {
	Message string `json:"stream"`
}

const baseContainerName = "ghcr.io/alexander-lindner/latex"

type RunCommand struct{}

var runCommand RunCommand
var cli *client.Client

func init() {
	_, err := parser.AddCommand("run",
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&runCommand)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}

	cli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic("Couldn't initialise the docker client.", err)
	}
}
func (x *RunCommand) Execute(args []string) error {
	mainConfig := options.Path + "/.latex"
	log.Println("Opening config file for  reading. Path:" + mainConfig)
	c, err := os.ReadFile(mainConfig)
	if err != nil {
		log.Fatal("Couldn't read the main config file", err)
	}
	config := configuration.ParseString(string(c))

	log.Println("Now pulling the base image: " + baseContainerName + ":base")
	err = pullImage(baseContainerName + ":base")
	if err != nil {
		log.Fatal("Couldn't pull the base docker image.", err)
	}
	image := config.GetString("docker.image", baseContainerName+":full")
	if image == "local" {
		log.Println("It is necessary to build the file before using it which is reasoned by choosing 'local' as the image in the config file")
		cwd, _ := os.Getwd()
		dir := filepath.Base(filepath.Dir(pathutil.CanonicalURLPath(cwd + "/" + options.Path)))

		hasher := sha1.New()
		hasher.Write([]byte(dir))
		sha := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
		sha = strings.ReplaceAll(sha, "+", "")

		imageName := strings.ToLower("textool-" + dir + "-" + sha + ":latest")

		err := buildImage(options.Path, imageName)
		if err == nil {
			image = imageName
		} else {
			log.Fatal("Building the image failed. Now fall back to full image")
		}
	}

	runImage(options.Path, image, config.GetString("fileName"))
	return nil
}
func runImage(basePath string, imageName string, outputName string) string {
	ctx := context.Background()

	u, err := user.Current()
	if err != nil {
		log.Panic("Couldn't retrieve current user", err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Panic("Couldn't retrieve current work directory", err)
	}
	log.Print("Creating the docker container")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Tty:   true,
		User:  u.Uid + ":" + u.Gid,
	}, &container.HostConfig{
		Binds:      []string{path + "/" + basePath + ":/data"},
		AutoRemove: true,
	}, nil, nil, "")

	if err != nil {
		log.Fatal("Couldn't create a container for running the latex commands", err)
	}
	log.Print("Starting the docker container")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal("Couldn't start the container for running the latex commands. ID of the container is "+resp.ID, err)
	}
	reader, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Timestamps: false,
		Follow:     true,
		Tail:       "40",
	})
	if err != nil {
		log.Fatal("Couldn't open logs of container. ", err)
	}

	err = printSimpleLog(reader)
	if err != nil {
		return ""
	}

	statusC, errC := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errC:
		if err != nil {
			log.Fatal("Couldn't wait for container. ", err)
		}
	case status := <-statusC:
		log.Println("Successfully run latex. Code:", status)
	}

	err = os.Rename(path+"/"+basePath+"/out/main.pdf", path+"/"+basePath+"/"+outputName)
	if err != nil {
		log.Fatal("Couldn't copy the file to the destination path. ", err)
	}
	return resp.ID
}

func pullImage(imageName string) error {
	ctx := context.Background()
	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {
			log.Fatal("Pulling the image failed.", err)
		}
	}(reader)
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		return err
	}
	return nil
}

func buildImage(path string, name string) error {
	ctx := context.Background()

	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{name},
		Remove:     true,
	}
	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Couldn't build the image.", err)
		}
	}(res.Body)

	err = printLog(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func printLog(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()

		msg := &DockerMessage{}
		err := json.Unmarshal([]byte(scanner.Text()), msg)
		if err != nil {
			return err
		}

		if msg.Message != "" {
			blue := color.FgBlue.Render
			fmt.Print(blue(msg.Message))
		}
	}

	errLine := &ErrorLine{}
	err := json.Unmarshal([]byte(lastLine), errLine)
	if err != nil {
		return err
	}
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func printSimpleLog(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()

		if lastLine != "" {
			blue := color.FgBlue.Render
			fmt.Println(blue(lastLine))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
