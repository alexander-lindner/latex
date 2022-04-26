package commands

import (
	"github.com/alexander-lindner/latex/textool/docker"
	"github.com/alexander-lindner/latex/textool/helper"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
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
	config, err := helper.GetConfig(options.Path)
	if err != nil {
		return x.executeWithNoConfigFile()
	} else {
		cli.SetFileName(config.GetString("fileName"))
		cli.SetTexFile(config.GetString("texFile"))

		log.Println("Now pulling the base image: " + baseContainerName + ":base")
		err = cli.PullImage(baseContainerName + ":base")
		if err != nil {
			log.Fatal("Couldn't pull the base docker image.", err)
		}
		Dockerfile := config.GetString("docker.file", "Dockerfile")
		image := cli.BuildLocalImage(Dockerfile)

		cli.RunImage(image)
		return nil
	}
}

func (x *RunCommand) executeWithNoConfigFile() error {
	log.Info("Falling back to no config mode.")

	listOfTExFiles := helper.ListFilesByExtension(options.Path, ".tex")
	if len(listOfTExFiles) == 1 {
		log.Info("Found one tex file, compiling it. ", listOfTExFiles[0])
		texFile := strings.Replace(listOfTExFiles[0], options.Path+"/", "", 1)
		cli.SetFileName(strings.Replace(texFile, ".tex", ".pdf", -1))
		cli.SetTexFile(texFile)
	} else {
		x.executeWithNoConfigFileAndMultipleTexFiles(listOfTExFiles)
	}
	Dockerfile := options.Path + "/Dockerfile"
	var image string
	if helper.PathExists(Dockerfile) {
		image = cli.BuildLocalImage("Dockerfile")
	} else {
		image = baseContainerName + ":full"
		err := cli.PullImage(image)
		if err != nil {
			log.Fatal("Couldn't pull the docker image.", err)
		}
	}
	cli.RunImage(image)

	return nil
}

func (x *RunCommand) executeWithNoConfigFileAndMultipleTexFiles(listOfTExFiles []string) {
	log.Info("Found more than one tex file, try to find the best one.")
	if listContains(listOfTExFiles, "main.tex") {
		log.Info("Found main.tex, compiling it.")
		cli.SetFileName("main.pdf")
		cli.SetTexFile("main.tex")
	} else {
		log.Info("Found no main.tex, compiling the first one which contains a \\documentclass definition.")
		var foundFiles []string
		for _, currentFile := range listOfTExFiles {
			log.Info("File: " + currentFile)
			data, err := ioutil.ReadFile(currentFile)
			if err == nil {
				foundDocumentClass := strings.Contains(string(data), "\\documentclass")
				if foundDocumentClass {
					foundFiles = append(foundFiles, currentFile)
				}
			} else {
				log.Error("Error reading file, skipping. ", err)
			}
		}

		if len(foundFiles) == 0 {
			log.Fatal("No tex file found which contains a \\documentclass definition. Specify your tex file with an minimal config file: https://textool.alindner.org/config/ ")

		} else {
			if len(foundFiles) == 1 {
				log.Info("Found one tex file, compiling it. ", foundFiles[0])
			} else {
				log.Info("Found more than one tex file: ", foundFiles)
				log.Info("using the first one: ", foundFiles[0])
				log.Info("It is recommended to specify your tex file with an minimal config file: https://textool.alindner.org/config/ ")
			}
			texFile := strings.Replace(foundFiles[0], options.Path+"/", "", 1)
			cli.SetFileName(strings.Replace(texFile, ".tex", ".pdf", -1))
			cli.SetTexFile(texFile)
		}
	}
}

func listContains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
