package githubapi

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"

	"Component-Manager/module"
)

func GetComponent(componentName string) (string, error) {
	COMPONENT_DIRECTORY := config.String("component_directory")
	SOURCE_COMPONENT_DIRECTORY := config.String("source.component_directory")

	packageManager, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			color.Redln("No package manager found. Please install one of the following package managers: pnpm, bun, yarn, npm.")
		}
		return "", errors.New("1")
	}
	color.Magentaf("Using:\t\t")
	color.Cyanln(packageManager)

	componentPath := filepath.Join(SOURCE_COMPONENT_DIRECTORY, componentName)

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

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
