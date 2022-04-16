package helper

import (
	"errors"
	"github.com/go-akka/configuration"
	log "github.com/sirupsen/logrus"
	"os"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		log.Panic("Couldn't fetch stats for "+path, err)
		return false
	}
}

func WriteConfig(config *configuration.Config, path string) {
	path = path + "/.latex"
	err := os.WriteFile(path, []byte(ConfigHeader+"\n"+config.String()), 0644)
	if err != nil {
		log.Panic("Couldn't write config to "+path, err)
	}

}
func GetConfig(path string) (config *configuration.Config) {
	mainConfig := path + "/.latex"
	log.Println("Opening config file for  reading. Path:" + mainConfig)
	c, err := os.ReadFile(mainConfig)
	if err != nil {
		log.Fatal("Couldn't read the main config file", err)
	}
	config = configuration.ParseString(string(c))
	return
}
