package commands

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {

}

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`
	Path    string `short:"p" long:"path" description:"the name of the directory, which should be created"`
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
			log.Fatal("During cli propagating an error raised.", err)
		default:
			log.Fatal("During cli propagating an unknown error raised.", err)
		}
	}
}
