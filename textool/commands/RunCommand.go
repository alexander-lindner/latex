package commands

import (
	"github.com/alexander-lindner/latex/textool/docker"
	"github.com/alexander-lindner/latex/textool/helper"
	log "github.com/sirupsen/logrus"
)

const baseContainerName = "ghcr.io/alexander-lindner/latex"

type RunCommand struct{}

var runCommand RunCommand

var cli docker.Client

func init() {
	_, err := parser.AddCommand("run",
		"Compiles a latex project",
		"Compiles a given latex project by first building the docker image and then use it to compile the project.",
		&runCommand)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
	cli = docker.New()
}
func (x *RunCommand) Execute(args []string) error {
	cli.SetBasePath(options.Path)
	config := helper.GetConfig(options.Path)

	cli.SetFileName(config.GetString("fileName"))
	cli.SetTexFile(config.GetString("texFile"))

	log.Println("Now pulling the base image: " + baseContainerName + ":base")
	err := cli.PullImage(baseContainerName + ":base")
	if err != nil {
		log.Fatal("Couldn't pull the base docker image.", err)
	}
	Dockerfile := config.GetString("docker.file", "Dockerfile")
	image := cli.BuildLocalImage(Dockerfile)

	cli.RunImage(image)
	return nil
}
