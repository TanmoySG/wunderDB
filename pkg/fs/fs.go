package fs

import (
	"errors"
	"os"
)

func CheckFileExists(filepath string) bool {
	_, error := os.Stat(filepath)
	if error == nil {
		return true
	} else if errors.Is(error, os.ErrNotExist) {
		return false
	}
	return true
}
