package githubapi

import (
	"errors"
	"path/filepath"
	"slices"
)

func FindComponent(componentName string) (string, error) {
	path := "Components/" + componentName

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

	file, directory, resp, err := githubClient.Repositories.GetContents(context, "LAZCO-STUDIO-LTD", "Component-Manager-Repo", path, nil)
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
