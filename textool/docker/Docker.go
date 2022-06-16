package docker

import (
	"bufio"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexander-lindner/latex/textool/helper"
	"github.com/coreos/etcd/pkg/pathutil"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/gookit/color"
	"github.com/moby/term"
	"github.com/radovskyb/watcher"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type Client struct {
	cli        *client.Client
	basePath   string
	outputName string
	texName    string
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
func (cl *Client) init() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Panic("Couldn't initialise the docker client.", err)
	}
	cl.cli = cli
}
func (cl *Client) RunImage(imageName string) string {
	return cl.RunImageWithCommand(
		imageName,
		"",
	)
}

func (cl *Client) PullImage(imageName string) error {
	if cl.IsImageAvailable(imageName) {
		log.Info("Image is already available. Not pulling.")
		return nil
	}
	ctx := context.Background()
	reader, err := cl.cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
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

func (cl *Client) BuildImage(path string, Dockerfile string, name string) error {
	ctx := context.Background()

	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: Dockerfile,
		Tags:       []string{name},
		Remove:     true,
	}
	res, err := cl.cli.ImageBuild(ctx, tar, opts)
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

func (cl *Client) SetBasePath(path string) {
	cl.basePath = path
}

func (cl *Client) BuildLocalImage(Dockerfile string) string {
	log.Println("It is necessary to build the file before using it")
	cwd, _ := os.Getwd()
	dir := filepath.Base(filepath.Dir(pathutil.CanonicalURLPath(cwd + "/" + cl.basePath)))

	hasher := sha1.New()
	hasher.Write([]byte(dir))
	sha := base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
	sha = strings.ReplaceAll(sha, "+", "")

	imageName := strings.ToLower("textool-" + dir + "-" + sha + ":latest")

	err := cl.BuildImage(cl.basePath, Dockerfile, imageName)
	if err == nil {
		return imageName
	} else {
		log.Error("Building the image failed. Now fall back to full image. ", err)
	}

	return baseContainerName + ":full"
}

func (cl *Client) RunImageWatch(imageName string) string {
	return cl.RunImageWithCommand(
		imageName,
		"watch",
	)
}
func (cl *Client) IsImageAvailable(imageName string) bool {
	ctx := context.Background()
	images, err := cl.cli.ImageList(ctx, types.ImageListOptions{All: true})
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
func (cl *Client) RunImageWithCommand(imageName string, command string) string {
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

	resp, err := cl.cli.ContainerCreate(
		ctx,
		config,
		&container.HostConfig{
			Binds:      []string{path + "/" + cl.basePath + ":/data"},
			AutoRemove: true,
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		log.Fatal("Couldn't create a container for running the latex commands", err)
	}

	sourcePdfName := strings.ReplaceAll(cl.texName, ".tex", ".pdf")
	originalPath := path + "/" + cl.basePath + "/out/" + sourcePdfName

	done := true

	if command == "watch" {
		if !helper.PathExists(originalPath) {
			log.Info("The final file does not exists. For watching it has to exists. Therefore, a normal build is executed before...")
			cl.RunImage(imageName)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {

				log.Println("Caught a [CTRL]+[C], stopping watch process ...")
				log.Debug(sig)

				d := 3 * time.Second
				if err := cl.cli.ContainerStop(ctx, resp.ID, &d); err != nil {
					log.Fatal("Couldn't stop the container. Error: ", err)
				}
				done = false
				log.Info("Stopped...")
			}
		}()

		go cl.watch(originalPath, &done, func(name string) {
			_, err = copy(originalPath, path+"/"+cl.basePath+"/"+cl.outputName)
			if err != nil {
				log.Fatal("Couldn't copy the file to the destination path. ", err)
			}
		}, imageName)
	}
	log.Print("Starting the docker container")
	if err := cl.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal("Couldn't start the container for running the latex commands. ID of the container is "+resp.ID, err)
	}
	reader, err := cl.cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
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

	statusC, errC := cl.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errC:
		if err != nil {
			log.Fatal("Couldn't wait for container. ", err)
		}
	case status := <-statusC:
		log.Println("Successfully run latex. Code:", status)
	}
	_, err = copy(originalPath, path+"/"+cl.basePath+"/"+cl.outputName)
	if err != nil {
		log.Fatal("Couldn't copy the file to the destination path. ", err)
	}
	return resp.ID
}

func (cl *Client) watch(path string, done *bool, callback func(name string), imageName string) {
	w := watcher.New()
	w.FilterOps(watcher.Write, watcher.Create)

	log.Info("Adding background watcher for changed files...")
	go func() {
		for *done {

			select {
			case event := <-w.Event:
				log.Info("Copy re-rendered file....")
				callback(event.Name())

			case err := <-w.Error:
				if err != nil {
					log.Println("error:", err)
				}
				go func() {
					log.Info("The final file was deleted. For watching it has to exists. Therefore, a normal build is executed before...")
					cl.RunImage(imageName)
					cl.watch(path, done, callback, imageName)
				}()
				w.Close()
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(path); err != nil {
		log.Fatal("Add failed:", err)
	}

	go func() {
		if err := w.Start(time.Second * 1); err != nil {
			log.Fatalln(err)
		}
	}()

}

func (cl *Client) SetFileName(fileName string) {
	cl.outputName = fileName
}

func (cl *Client) SetTexFile(texName string) {
	cl.texName = texName
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
