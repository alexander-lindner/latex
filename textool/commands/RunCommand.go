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
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&runCommand)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
	cli = docker.New()
}
func (x *RunCommand) Execute(args []string) error {
	cli.SetBasePath(options.Path)
	config := helper.GetConfig(options.Path)

	log.Println("Now pulling the base image: " + baseContainerName + ":base")
	err := cli.PullImage(baseContainerName + ":base")
	if err != nil {
		log.Fatal("Couldn't pull the base docker image.", err)
	}
	image := config.GetString("docker.image", baseContainerName+":full")
	if image == "local" {
		image = cli.BuildLocalImage()
	} else {
		err = cli.PullImage(baseContainerName + ":full")
		if err != nil {
			log.Fatal("Couldn't pull the full docker image.", err)
		}
	}

	cli.RunImage(options.Path, image, config.GetString("fileName"))
	return nil
}
