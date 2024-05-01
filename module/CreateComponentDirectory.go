package module

import (
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"
)

func CreateComponentDirectory(projectPath string) (string, error) {
	COMPONENT_DIRECTORY := config.String("component_directory")

	path := filepath.Join(projectPath, COMPONENT_DIRECTORY)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		return path, err
	} else {
		color.Yellowp("Warning: Component directory ")
		color.Cyanp(COMPONENT_DIRECTORY)
		color.Yellowln(" already exists. Using existing directory.")
		return path, nil
	}
}
