package docker

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
	"github.com/gookit/color"
	"github.com/moby/term"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Client struct {
	cli      *client.Client
	basePath string
}

const baseContainerName = "ghcr.io/alexander-lindner/latex"

var c = color.FgLightBlue

func init() {

}
func New() Client {
	var cli Client
	cli.init()
	return cli
}
func (this *Client) init() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic("Couldn't initialise the docker client.", err)
	}
	this.cli = cli
}
func (this *Client) RunImage(basePath string, imageName string, outputName string) string {
	return this.RunImageWithCommand(
		basePath,
		imageName,
		outputName,
		"",
	)
}

func (this *Client) PullImage(imageName string) error {
	if this.IsImageAvailable(imageName) {
		log.Info("Image is already available. Not pulling.")
		return nil
	}
	ctx := context.Background()
	reader, err := this.cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
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
	_, err = color.Set(c)
	if err != nil {
		return err
	}
	err = jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)
	_, err = color.Reset()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *Client) BuildImage(path string, name string) error {
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
	res, err := this.cli.ImageBuild(ctx, tar, opts)
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

func (this *Client) SetBasePath(path string) {
	this.basePath = path
}

func (this *Client) BuildLocalImage() string {
	log.Println("It is necessary to build the file before using it which is reasoned by choosing 'local' as the image in the config file")
	cwd, _ := os.Getwd()
	dir := filepath.Base(filepath.Dir(pathutil.CanonicalURLPath(cwd + "/" + this.basePath)))

	hasher := sha1.New()
	hasher.Write([]byte(dir))
	sha := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
	sha = strings.ReplaceAll(sha, "+", "")

	imageName := strings.ToLower("textool-" + dir + "-" + sha + ":latest")

	err := this.BuildImage(this.basePath, imageName)
	if err == nil {
		var image = imageName
		return image
	} else {
		log.Fatal("Building the image failed. Now fall back to full image. ", err)
	}

	return baseContainerName + ":full"
}

func (this *Client) RunImageWatch(basePath string, imageName string, outputName string) string {
	return this.RunImageWithCommand(
		basePath,
		imageName,
		outputName,
		"watch",
	)
}
func (this *Client) IsImageAvailable(imageName string) bool {
	ctx := context.Background()
	images, err := this.cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		log.Fatal("Couldn't list docker images. ", err)
	}
	for _, img := range images {
		if strings.Contains(img.RepoTags[0], imageName) {
			return true
		}
	}
	return false
}
func (this *Client) RunImageWithCommand(basePath string, imageName string, outputName string, command string) string {
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
	var config = &container.Config{
		Image:        imageName,
		Tty:          true,
		User:         u.Uid + ":" + u.Gid,
		AttachStderr: true,
		AttachStdout: true,
		AttachStdin:  true,
		OpenStdin:    true,
		StdinOnce:    true,
	}
	if command != "" {
		config.Cmd = []string{command}
	}

	resp, err := this.cli.ContainerCreate(
		ctx,
		config,
		&container.HostConfig{
			Binds:      []string{path + "/" + basePath + ":/data"},
			AutoRemove: true,
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		log.Fatal("Couldn't create a container for running the latex commands", err)
	}
	log.Print("Starting the docker container")
	if err := this.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal("Couldn't start the container for running the latex commands. ID of the container is "+resp.ID, err)
	}
	reader, err := this.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
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

	statusC, errC := this.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errC:
		if err != nil {
			log.Fatal("Couldn't wait for container. ", err)
		}
	case status := <-statusC:
		log.Println("Successfully run latex. Code:", status)
	}

	_, err = copy(path+"/"+basePath+"/out/main.pdf", path+"/"+basePath+"/"+outputName)
	if err != nil {
		log.Fatal("Couldn't copy the file to the destination path. ", err)
	}
	return resp.ID
}
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
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
			blue := c.Render
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
			blue := c.Render
			fmt.Println(blue(lastLine))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

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