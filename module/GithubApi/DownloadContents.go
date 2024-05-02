package githubapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v61/github"
	"github.com/gookit/color"
	"github.com/gookit/config/v2"
)

func DownloadFile(fileContent *github.RepositoryContent, componentPath string) error {
	COMPONENT_DIRECTORY := config.String("component_directory")

	componentPathParts := strings.Split(componentPath, "/")
	componentName := strings.Join(componentPathParts[1:], "/")

	filePath := filepath.Join(COMPONENT_DIRECTORY, componentName, fileContent.GetName())

	downloadURL := fileContent.GetDownloadURL()

	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.New("failed to create directory: " + dirPath)
	}

	resp, err := http.Get(downloadURL)
	if err != nil {
		return errors.New("failed to download the file: " + downloadURL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("failed to download the file: " + downloadURL)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return errors.New("failed to create file: " + filePath)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.New("failed to write file: " + filePath)
	}

	color.Magentap("Downloaded:\t")
	color.Cyanln(filepath.Join(componentName, fileContent.GetName()))

	return nil
}

func DownloadComponent(githubClient *github.Client, context context.Context, componentPath string) error {
	OWNER := config.String("source.owner")
	REPO := config.String("source.repo")

	_, directoryContents, _, err := githubClient.Repositories.GetContents(context, OWNER, REPO, componentPath, nil)
	if err != nil {
		return err
	}

	for _, content := range directoryContents {
		if *content.Type == "dir" {
			err := DownloadComponent(githubClient, context, *content.Path)
			if err != nil {
				return err
			}
		} else if *content.Type == "file" {
			err := DownloadFile(content, componentPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
