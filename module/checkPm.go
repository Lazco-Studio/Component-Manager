package module

import (
	"os/exec"

	"github.com/gookit/color"
)

func CheckPm() string {
	for _, bin := range []string{"pnpm", "bun", "yarn", "npm"} {
		if _, err := exec.LookPath(bin); err == nil {
			color.Greenln("Package manager found:", bin)
			return bin
		}
	}
	color.Redln("No package manager found, please install one of pnpm, bun, yarn, npm")
	return ""
}
