package githubapi

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

func GetComponent(path string) (string, error) {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		return "", errors.New("please set the GITHUB_TOKEN environment variable")
	}

	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(context, ts)
	githubClient := github.NewClient(tc)

	file, directory, resp, err := githubClient.Repositories.GetContents(context, "LAZCO-STUDIO-LTD", "Component-Manager-Repo", path, nil)
	if err != nil {
		if resp.StatusCode != 200 {
			return "", errors.New("not found")
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
