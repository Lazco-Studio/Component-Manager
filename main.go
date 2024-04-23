package main

import (
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/LAZCO-STUDIO-LTD/Component-Manager/command"
)

var GITHUB_TOKEN string

func main() {
	os.Setenv("GITHUB_TOKEN", GITHUB_TOKEN)

	app := &cli.App{
		Name:     "component-manager",
		HelpName: "cm",
		Usage:    "an component manager",
		Version:  "v0.2.5",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version of ui",
				Action:  command.Version,
			},
			{
				Name:   "init",
				Usage:  "initialize a new project",
				Action: command.Init,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a new component",
				Action:  command.Add,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		var errorString = err.Error()

		if errorNumber, err := strconv.Atoi(errorString); err == nil {
			os.Exit(errorNumber)
		}

		color.Redln(errorString)
		os.Exit(1)
	}
}
