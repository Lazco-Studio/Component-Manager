package module

import (
	"errors"
	"os"
	"path/filepath"
)

func GetPath(path string) (string, error) {
	if path == "" {
		directory, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return directory, nil
	} else {
		fileInfo, err := os.Stat(path)
		if os.IsNotExist(err) {
			return "", errors.New("path does not exist")
		} else if err != nil {
			return "", err
		}

		if fileInfo.IsDir() {
			path, err = filepath.Abs(path)
			if err != nil {
				return "", err
			}
			return path, nil
		}
		return "", errors.New("path is not a directory")
	}
}
