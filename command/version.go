package command

import (
	"github.com/urfave/cli/v2"

	"Component-Manager/module"
)

func Version(ctx *cli.Context) error {
	cli.ShowVersion(ctx)
	module.CheckRemoteVersion(ctx, false)
	return nil
}
