package githubapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v61/github"
	"github.com/gookit/color"
	"github.com/gookit/config/v2"

	"Component-Manager/module"
)

// downloadFile downloads a file from its GitHub download URL and saves it to the specified local path.
func downloadFile(fileContent *github.RepositoryContent, componentPath string) error {
	COMPONENT_DIRECTORY := config.String("component_directory")

	// remove first path in componentPath
	componentPathParts := strings.Split(componentPath, "/")
	componentName := strings.Join(componentPathParts[1:], "/")

	filePath := filepath.Join(COMPONENT_DIRECTORY, componentName, fileContent.GetName())

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

	color.Magentap("Downloaded:\t")
	color.Cyanln(filepath.Join(componentName, fileContent.GetName()))

	return nil
}

// downloadContents recursively downloads files from a given repository path.
func downloadContents(githubClient *github.Client, context context.Context, componentPath string) error {
	OWNER := config.String("source.owner")
	REPO := config.String("source.repo")

	_, directoryContents, _, err := githubClient.Repositories.GetContents(context, OWNER, REPO, componentPath, nil)
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
	COMPONENT_DIRECTORY := config.String("component_directory")
	SOURCE_COMPONENT_DIRECTORY := config.String("source.component_directory")

	packageManager, err := module.CheckPm()
	if err != nil {
		switch err.Error() {
		case "no package manager found":
			color.Redln("No package manager found. Please install one of the following package managers: pnpm, bun, yarn, npm.")
		}
		return "", errors.New("1")
	}
	color.Magentaf("Using:\t\t")
	color.Cyanln(packageManager)

	componentPath := filepath.Join(SOURCE_COMPONENT_DIRECTORY, componentName)

	githubClient, context, err := GithubClient()
	if err != nil {
		return "", err
	}

	err = downloadContents(githubClient, context, componentPath)
	if err != nil {
		return "", err
	}

	color.Yellowln("Installing dependencies...")
	module.FullWidthMessage("installation log start", color.Gray)
	cmd := exec.Command(packageManager, "install")
	cmd.Dir = filepath.Join(COMPONENT_DIRECTORY, componentName)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	module.FullWidthMessage("installation log end", color.Gray)

	return componentName, nil
}
