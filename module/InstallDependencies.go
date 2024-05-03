package module

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"
)

func InstallDependencies(packageManager string, componentName string) (string, error) {
	COMPONENT_DIRECTORY := config.String("component_directory")

	color.Yellowln("\nInstalling dependencies...")
	FullWidthMessage("installation log start", color.Gray)

	// os.Setenv("NODE_ENV", "production")
	var cmd *exec.Cmd
	switch packageManager {
	case "npm":
		cmd = exec.Command(packageManager, "install", "--production", "--legacy-peer-deps")
	case "yarn":
		cmd = exec.Command(packageManager, "workspaces", "focus", "--production")
	default:
		cmd = exec.Command(packageManager, "install", "--production")
	}

	cmd.Dir = filepath.Join(COMPONENT_DIRECTORY, componentName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	FullWidthMessage("installation log end", color.Gray)

	return componentName, nil
}
