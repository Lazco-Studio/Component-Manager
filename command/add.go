package command

import (
	"errors"
	"os"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/urfave/cli/v2"

	"Component-Manager/module/GithubApi"
)

func Add(ctx *cli.Context) error {
	COMPONENT_DIRECTORY := config.String("component_directory")

	componentName := ctx.Args().Get(0)

	if _, err := os.Stat(COMPONENT_DIRECTORY); errors.Is(err, os.ErrNotExist) {
		return errors.New("project isn't initialized yet. Please run 'cm init' first")
	}

	_, err := githubapi.FindComponent(componentName)
	if err != nil {
		switch err.Error() {
		case "not found":
			fallthrough
		case "not a directory":
			fallthrough
		case "not a component":
			return errors.New("component " + componentName + " not found")

		default:
			return err
		}
	}

	fileContent, err := githubapi.GetComponent(componentName)
	if err != nil {
		return err
	}

	color.Greenp("Successfully added component: ")
	color.Cyanln(fileContent)

	return nil
}
