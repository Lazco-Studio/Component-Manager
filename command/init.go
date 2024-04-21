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
	color.Greenf("Using:\t\t")
	color.Cyanln(packageManager)

	projectPath, err := module.GetPath(ctx.Args().Get(0))
	if err != nil {
		switch err.Error() {
		case "path does not exist":
			color.Redln("Path does not exist.")
		case "path is not a directory":
			color.Redln("Path is not a directory.")
		}
		return errors.New("1")
	}

	color.Greenf("Project Path:\t")
	color.Cyanln(projectPath)

	componentPath, err := module.CreateComponentDirectory(projectPath)
	if err != nil {
		color.Redln(err.Error())
		return errors.New("1")
	}

	color.Greenf("Component path:\t")
	color.Cyanln(componentPath)

	// module.FullWidthMessage("installation log start", color.Gray)
	// err = module.InstallNodePackage(projectPath, packageManager, false, "fs")
	// module.FullWidthMessage("installation log end", color.Gray)
	// if err != nil {
	// 	color.Redln(err.Error())
	// 	return errors.New("1")
	// }

	return nil
}
