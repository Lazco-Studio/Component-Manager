package main

import (
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/LAZCO-STUDIO-LTD/Component-Manager/command"
)

func main() {
	app := &cli.App{
		Name:     "component-manager",
		HelpName: "cm",
		Usage:    "an component manager",
		Version:  "v0.0.0",
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
		},
	}

	if err := app.Run(os.Args); err != nil {
		var errorString = err.Error()

		if errorNumber, err := strconv.Atoi(errorString); err == nil {
			os.Exit(errorNumber)
		}

		log.Fatal(err)
		os.Exit(1)
	}
}
