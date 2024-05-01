package main

import (
	"embed"
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/urfave/cli/v2"

	"Component-Manager/command"
)

//go:embed .cmrc.json
var configFile embed.FS

var GITHUB_TOKEN string

func main() {
	os.Setenv("GITHUB_TOKEN", GITHUB_TOKEN)
	configFileContent, err := configFile.ReadFile(".cmrc.json")
	if err != nil {
		color.Redln(err)
		os.Exit(1)
	}

	err = config.LoadSources("json", configFileContent)
	if err != nil {
		color.Redln(err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:     "Component-Manager",
		HelpName: "cm",
		Usage:    "A tool for managing JS/TS components and modules.",
		Version:  "v1.2.1",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version of cm",
				Action:  command.Version,
			},
			{
				Name:   "init",
				Usage:  "initialize a new project",
				Action: command.Init,
			},
			{
				Name:    "add",
				Aliases: []string{"a", "get", "download"},
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
