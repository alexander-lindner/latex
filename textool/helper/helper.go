package helper

import (
	"errors"
	"github.com/go-akka/configuration"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
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
func GetConfig(path string) (config *configuration.Config, er error) {
	mainConfig := path + "/.latex"
	log.Println("Opening config file for  reading. Path:" + mainConfig)
	c, err := os.ReadFile(mainConfig)
	if err != nil {
		log.Error("Couldn't read the main config file. ", err)
		return nil, err
	}
	config = configuration.ParseString(string(c))
	return
}

func ListFilesByExtension(path string, extension string) []string {
	var listOfFiles []string
	err := filepath.WalkDir(path, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == extension {
			listOfFiles = append(listOfFiles, s)
		}
		return nil
	})
	if err != nil {
		log.Error("Couldn't list files in "+path+". ", err)
		return nil
	}
	return listOfFiles
}
