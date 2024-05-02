package githubapi

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"

	"Component-Manager/module"
)

func GetComponent(componentName string) (string, error) {
	COMPONENT_DIRECTORY := config.String("component_directory")
	SOURCE_COMPONENT_DIRECTORY := config.String("source.component_directory")
	PACKAGE_MANAGERS := config.Strings("package_managers")

	packageManager, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			return "", errors.New("no package manager found, please install one of the following package managers: " + strings.Join(PACKAGE_MANAGERS, ", ") + ".")
		}
		return "", err
	}
	color.Magentaf("Using package manager: ")
	color.Cyanln(packageManager)

	componentPath := filepath.Join(SOURCE_COMPONENT_DIRECTORY, componentName)

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

	color.Magentaln("Downloaded:")
	err = DownloadComponent(githubClient, context, componentPath)
	if err != nil {
		return "", err
	}

	color.Yellowln("Installing dependencies...")
	module.FullWidthMessage("installation log start", color.Gray)
	cmd := exec.Command(packageManager, "install")
	cmd.Dir = filepath.Join(COMPONENT_DIRECTORY, componentName)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	module.FullWidthMessage("installation log end", color.Gray)

	return componentName, nil
}
