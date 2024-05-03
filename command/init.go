package command

import (
	"errors"
	"strings"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/gookit/config/v2"

	"Component-Manager/module"
)

func Init(ctx *cli.Context) error {
	PACKAGE_MANAGERS := config.Strings("package_managers")

	_, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			return errors.New("no package manager found, please install one of the following package managers: " + strings.Join(PACKAGE_MANAGERS, ", ") + ".")
		}
		return err
	}
	color.Magentaln("Found package manager:")
	for _, pm := range PACKAGE_MANAGERS {
		color.Grayp(" - ")
		color.Cyanln(pm)
	}

	projectPath, err := module.GetPath(ctx.Args().Get(0))
	if err != nil {
		switch err.Error() {
		case "path does not exist":
			return errors.New("specified path does not exist.")
		case "path is not a directory":
			return errors.New("specified path is not a directory.")
		}
		return err
	}

	componentPath, err := module.CreateComponentDirectory(projectPath)
	if err != nil {
		return err
	}

	color.Magentaf("Component directory: ")
	color.Cyanln(componentPath)

	return nil
}
