package module

import (
	"os"
	"path/filepath"
)

func CreateComponentDirectory(projectPath string) (string, error) {
	path := filepath.Join(projectPath, "lazco_components")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		return path, err
	}

	return path, nil
}
