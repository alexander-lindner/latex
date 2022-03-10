package commands

import (
	"github.com/geomyidia/flagswrap"
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
		wrappedErr := flagswrap.WrapError(err)
		switch {
		case wrappedErr.IsHelp():
			os.Exit(0)
		// Self-documenting go-flags errors:
		case wrappedErr.IsVerbose():
			os.Exit(1)
		// go-flags errors that need more context:
		case wrappedErr.IsSilent():
			// TODO: if you see duplicate error messages here, then
			// you just need to move the error in question from the
			// goFlagsSilentErrors to the goFlagsVerboseErrors map
			// in ./errors.go -- and then submit a PR!
			log.Fatal("During cli propagating an error raised.", wrappedErr)
			os.Exit(1)
		default:
			// TODO: anything here might justify a PR ...
			log.Fatal("During cli propagating an unknown error raised.", wrappedErr)
			os.Exit(1)
		}
	}
}
