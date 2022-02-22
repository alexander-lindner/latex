package helper

import (
	"errors"
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
