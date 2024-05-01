package command

import (
	"github.com/urfave/cli/v2"
)

func Version(ctx *cli.Context) error {
	cli.ShowVersion(ctx)
	return nil
}
