package githubapi

import (
	"errors"
	"path/filepath"
	"slices"

	"github.com/gookit/config/v2"
)

func FindComponent(componentName string) (string, error) {
	SOURCE_COMPONENT_DIRECTORY := config.String("source.component_directory")
	OWNER := config.String("source.owner")
	REPO := config.String("source.repo")

	path := filepath.Join(SOURCE_COMPONENT_DIRECTORY, componentName)

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

	file, directory, resp, err := githubClient.Repositories.GetContents(context, OWNER, REPO, path, nil)
	if err != nil {
		if resp.StatusCode == 404 {
			return "", errors.New("not found")
		}
		if (resp.StatusCode == 401) || (resp.StatusCode == 403) {
			return "", errors.New("invalid github token")
		}
		return "", err
	}

	if file != nil {
		return "", errors.New("not a directory")
	}

	if directory != nil {
		var fileList []string

		for _, file := range directory {
			fileList = append(fileList, *file.Name)
		}

		checkList := []string{"types", "index.ts", "package.json"}
		for _, checkFile := range checkList {
			if !slices.Contains(fileList, checkFile) {
				return "", errors.New("not a component")
			}
		}

		return filepath.Clean(*directory[0].Path + "/.."), nil
	}

	return "", nil
}
