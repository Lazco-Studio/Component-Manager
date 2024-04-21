package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/LAZCO-STUDIO-LTD/Component-Manager/module"
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
				Action: func(ctx *cli.Context) error {
					fmt.Println(ctx.App.Version)
					return nil
				},
			},
			{
				Name:  "init",
				Usage: "initialize a new project",
				Action: func(ctx *cli.Context) error {
					pm := module.CheckPm()
					fmt.Println(pm)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
