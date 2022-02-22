package commands

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"os"
	"os/user"
)

const baseContainerName = "ghcr.io/alexander-lindner/latex"

type RunCommand struct{}

var runCommand RunCommand

func (x *RunCommand) Execute(args []string) error {
	runImage(options.Path)
	return nil
}

func init() {
	_, err := parser.AddCommand("run",
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&runCommand)
	if err != nil {
		return
	}
}

func runImage(basePath string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	u, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()
	var path, _ = os.Getwd()
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: baseContainerName + ":full",
		Tty:   true,
		User:  u.Uid + ":" + u.Gid,
	}, &container.HostConfig{
		Binds:      []string{path + "/" + basePath + ":/data"},
		AutoRemove: true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	return resp.ID
}
