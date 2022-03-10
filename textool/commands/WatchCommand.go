package commands

import (
	"github.com/alexander-lindner/latex/textool/docker"
	"github.com/alexander-lindner/latex/textool/helper"
	log "github.com/sirupsen/logrus"
)

type WatchCommand struct{}

var watchCommand WatchCommand

func init() {
	_, err := parser.AddCommand("watch",
		"Initialise a latex project directory",
		"Creates a directory and adds a minimal Latex template to this directory",
		&watchCommand)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
	cli = docker.New()
}
func (x *WatchCommand) Execute(args []string) error {
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

	cli.RunImageWatch(options.Path, image, config.GetString("fileName"))
	return nil
}
