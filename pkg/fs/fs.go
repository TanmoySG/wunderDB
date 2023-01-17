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

func CreateFile(filepath string) error {
	err := os.WriteFile(filepath, []byte{}, 0740)
	if err != nil {
		return err
	}
	return nil
}

func CreateDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	return err
}

func WriteToFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0740)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filepath string) ([]byte, error) {
	dataBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
