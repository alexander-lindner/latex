package main

import (
	"github.com/alexander-lindner/latex/textool/commands"
	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetFlags(0)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(false)

	commands.EvalCli()
}
