package module

import (
	"errors"
	"os/exec"
)

func CheckPm() (string, error) {
	var checkList = []string{"pnpm", "bun", "yarn", "npm"}

	for _, bin := range checkList {
		if _, err := exec.LookPath(bin); err == nil {
			return bin, nil
		}
	}

	return "", errors.New("no package manager found")
}
