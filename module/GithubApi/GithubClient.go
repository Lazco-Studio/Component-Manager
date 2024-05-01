package githubapi

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

func GithubClient() (*github.Client, context.Context, error) {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		return nil, nil, errors.New("please set the GITHUB_TOKEN environment variable")
	}

	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(context, ts)
	githubClient := github.NewClient(tc)

	return githubClient, context, nil
}
