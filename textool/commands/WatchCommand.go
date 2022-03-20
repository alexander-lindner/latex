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
		"Build and watches a latex project",
		"Compiles a latex project and afterwards, watches for changes which triggers a recompilation",
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
	Dockerfile := config.GetString("docker.file", "Dockerfile")

	image := cli.BuildLocalImage(Dockerfile)

	cli.RunImageWatch(options.Path, image, config.GetString("fileName"), config.GetString("texFile"))
	return nil
}
