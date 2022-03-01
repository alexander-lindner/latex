package commands

import (
	"github.com/alexander-lindner/latex/textool/helper"
	log "github.com/sirupsen/logrus"
	"os"
)

type CleanCommand struct {
}

var cleanCommand CleanCommand
var toRemove = [...]string{"out", ".texlive2020", ".texlive2021", ".texlive2022"}

func (x *CleanCommand) Execute(args []string) error {
	log.Println("Cleaning all temporary files...")
	for _, path := range toRemove {
		realPath := options.Path + "/" + path
		if helper.PathExists(realPath) {
			err := os.RemoveAll(realPath)
			if err != nil {
				log.Panic("Couldn't remove necessary path "+path, err)
			}
		}
	}
	return nil
}

func init() {
	_, err := parser.AddCommand("clean",
		"Cleans the working directory",
		"Removes all temporarily created files and directories",
		&cleanCommand,
	)
	if err != nil {
		log.Panic("Building the command parameter went wrong.", err)
	}
}
