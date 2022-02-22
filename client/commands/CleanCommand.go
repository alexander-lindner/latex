package commands

import (
	"github.com/alexander-lindner/latex/client/latex/helper"
	"os"
)

type CleanCommand struct {
}

var cleanCommand CleanCommand
var toRemove = [...]string{"out", ".texlive2020", ".texlive2021", ".texlive2022"}

func (x *CleanCommand) Execute(args []string) error {
	for _, path := range toRemove {
		realPath := options.Path + "/" + path
		helper.Exists(realPath, func() {
			err := os.RemoveAll(realPath)
			if err != nil {
				panic(err)
			}
		}, func() {

		})
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
		return
	}
}
