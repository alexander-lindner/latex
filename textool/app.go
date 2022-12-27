package main

import (
	"github.com/alexander-lindner/latex/textool/commands"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	log "github.com/sirupsen/logrus"
)

const version = "2.1.7"

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "alexander-lindner/latex")
	if err != nil {
		log.Error("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		log.Debug("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}
func main() {
	//log.SetFlags(0)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(false)

	doSelfUpdate()

	commands.EvalCli()
}
