package command

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"Component-Manager/module"
)

func DownloadScript() string {
	tempDir := os.TempDir()

	scriptDir := module.Download("https://short.on-cloud.tw/cm-install-script", filepath.Join(tempDir, "cm-install-script"))
	os.Chmod(scriptDir, 0755)

	return scriptDir
}

func Update(ctx *cli.Context) error {
	scriptDir := DownloadScript()
	cmd := exec.Command(scriptDir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.Remove(scriptDir)

	return nil
}
