package module

import (
	"os"
	"os/exec"
)

func InstallNodePackage(absolutePath string, packageManager string, isDevDependency bool, packageName string) error {
	cmd := exec.Command(packageManager, "install", packageName, DevDependencyArg(isDevDependency))
	cmd.Dir = absolutePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func DevDependencyArg(devDependency bool) string {
	if devDependency {
		return "-D"
	}
	return ""
}
