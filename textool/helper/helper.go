package helper

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func Exists(path string, yes func(), no func()) {
	if _, err := os.Stat(path); err == nil {
		yes()
	} else if errors.Is(err, os.ErrNotExist) {
		no()
	} else {
		no()
	}
}

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
