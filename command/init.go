package command

import (
	"errors"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/LAZCO-STUDIO-LTD/Component-Manager/module"
)

func Init(ctx *cli.Context) error {
	packageManager, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			color.Redln("No package manager found. Please install one of the following package managers: pnpm, bun, yarn, npm.")
		}
		return errors.New("1")
	}
	color.Magentaf("Using:\t\t")
	color.Cyanln(packageManager)

	projectPath, err := module.GetPath(ctx.Args().Get(0))
	if err != nil {
		switch err.Error() {
		case "path does not exist":
			color.Redln("Specified path does not exist.")
		case "path is not a directory":
			color.Redln("Specified path is not a directory.")
		}
		return errors.New("1")
	}

	color.Magentaf("Project Path:\t")
	color.Cyanln(projectPath)

	componentPath, err := module.CreateComponentDirectory(projectPath)
	if err != nil {
		color.Redln(err.Error())
		return errors.New("1")
	}

	color.Magentaf("Component path:\t")
	color.Cyanln(componentPath)

	return nil
}
