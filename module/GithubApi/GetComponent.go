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
)

// downloadFile downloads a file from its GitHub download URL and saves it to the specified local path.
func downloadFile(fileContent *github.RepositoryContent, componentPath string) error {
	// remove first path in componentPath
	componentPathParts := strings.Split(componentPath, "/")
	componentName := strings.Join(componentPathParts[1:], "/")

	filePath := filepath.Join("lazco_components", componentName, fileContent.GetName())

	// Get the download URL for the file.
	downloadURL := fileContent.GetDownloadURL()

	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return errors.New("failed to create directory: " + dirPath)
	}

	// Perform an HTTP GET request to download the file content.
	resp, err := http.Get(downloadURL)
	if err != nil {
		return errors.New("failed to download the file: " + downloadURL)
	}
	defer resp.Body.Close()

	// Ensure the request was successful.
	if resp.StatusCode != 200 {
		return errors.New("failed to download the file: " + downloadURL)
	}

	// Create a local file where the content will be saved.
	out, err := os.Create(filePath)
	if err != nil {
		return errors.New("failed to create file: " + filePath)
	}
	defer out.Close()

	// Copy the downloaded content to the local file.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.New("failed to write file: " + filePath)
	}

	color.Magentap("Downloaded: ")
	color.Cyanln(filepath.Join(componentName, fileContent.GetName()))

	return nil
}

// downloadContents recursively downloads files from a given repository path.
func downloadContents(githubClient *github.Client, context context.Context, componentPath string) error {
	owner := "LAZCO-STUDIO-LTD"
	repo := "Component-Manager-Repo"

	_, directoryContents, _, err := githubClient.Repositories.GetContents(context, owner, repo, componentPath, nil)
	if err != nil {
		return err
	}

	for _, content := range directoryContents {
		if *content.Type == "dir" {
			// Recursive call for subdirectories
			err := downloadContents(githubClient, context, *content.Path)
			if err != nil {
				return err
			}
		} else if *content.Type == "file" {
			// Download file
			err := downloadFile(content, componentPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetComponent(componentName string) (string, error) {
	componentPath := "Components/" + componentName

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

	err = downloadContents(githubClient, context, componentPath)
	if err != nil {
		return "", err
	}

	return componentName, nil
}
