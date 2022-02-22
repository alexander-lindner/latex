package commands

import (
	"github.com/jessevdk/go-flags"
	"os"
)

func init() {

}

type Options struct {
	// Example of verbosity with level
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`

	Path string `short:"p" long:"path" description:"the name of the directory, which should be created"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func EvalCli() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
