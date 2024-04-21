package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Version(ctx *cli.Context) error {
	fmt.Println(ctx.App.Version)
	return nil
}
