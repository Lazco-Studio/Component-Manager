package command

import (
	"errors"
	"os"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/LAZCO-STUDIO-LTD/Component-Manager/module/GithubApi"
)

func Add(ctx *cli.Context) error {
	componentName := ctx.Args().Get(0)

	if _, err := os.Stat("lazco_components"); errors.Is(err, os.ErrNotExist) {
		color.Redln("Project isn't initialized yet. Please run 'cm init' first.")
		return errors.New("1")
	}

	path, err := githubapi.GetComponent("Components/" + componentName)
	if err != nil {
		switch err.Error() {
		case "not found":
			fallthrough
		case "not a directory":
			fallthrough
		case "not a component":
			color.Redln("Component " + componentName + " not found.")
			return errors.New("1")

		default:
			return err
		}
	}

	color.Magentaf("URL: ")
	color.Cyanln("https://raw.githubusercontent.com/LAZCO-STUDIO-LTD/Component-Manager-Repo/main/" + path)

	return nil
}
