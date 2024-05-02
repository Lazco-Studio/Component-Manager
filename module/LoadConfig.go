package module

import (
	"errors"
	"os"

	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/urfave/cli/v2"
)

func LoadConfig(configPath string, officialConfigBytes []byte) (string, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = config.LoadSources("json", officialConfigBytes)
		if err != nil {
			return "", errors.New("failed to load official config file")
		}
		return ".cmrc.official.json", nil
	}

	err := config.LoadFiles(configPath)
	if err != nil {
		return "", errors.New("failed to load config file")
	}

	packageManager := config.Strings("package_managers")
	if len(packageManager) == 0 {
		return "", errors.New("no package manager found in config file")
	}

	componentDirectory := config.String("component_directory")
	if componentDirectory == "" {
		return "", errors.New("no component directory found in config file")
	}

	source_owner := config.String("source.owner")
	if source_owner == "" {
		return "", errors.New("no source owner found in config file")
	} else if !GithubNamingRule(source_owner) {
		return "", errors.New("invalid source.owner in config file")
	}

	source_repo := config.String("source.repo")
	if source_repo == "" {
		return "", errors.New("no source repo found in config file")
	} else if !GithubNamingRule(source_repo) {
		return "", errors.New("invalid source.repo in config file")
	}

	source_componentDirectory := config.String("source.component_directory")
	if source_componentDirectory == "" {
		return "", errors.New("no source component directory found")
	}

	color.Magentap("Using custom config: ")
	color.Cyanln(configPath)

	return configPath, nil
}

func LoadAppConfig(ctx *cli.Context, officialConfigBytes []byte) error {
	_, err := LoadConfig(".cmrc.json", officialConfigBytes)
	return err
}
