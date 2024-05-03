package githubapi

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"

	"Component-Manager/module"
)

func GetComponent(componentName string) (string, error) {
	SOURCE_COMPONENT_DIRECTORY := config.String("source.component_directory")
	PACKAGE_MANAGERS := config.Strings("package_managers")

	packageManager, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			return "", errors.New("no package manager found, please install one of the following package managers: " + strings.Join(PACKAGE_MANAGERS, ", "))
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

	color.Magentaln("\nDownloaded:")
	err = DownloadComponent(githubClient, context, componentPath)
	if err != nil {
		return "", err
	}

	componentName, err = module.InstallDependencies(packageManager, componentName)
	if err != nil {
		return "", err
	}

	return componentName, nil
}
