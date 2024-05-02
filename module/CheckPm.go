package module

import (
	"errors"
	"os/exec"

	"github.com/gookit/config/v2"
)

func CheckPm() (string, error) {
	PACKAGE_MANAGERS := config.Strings("package_managers")

	for _, bin := range PACKAGE_MANAGERS {
		if _, err := exec.LookPath(bin); err == nil {
			return bin, nil
		}
	}

	return "", errors.New("no package manager found")
}
