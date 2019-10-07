package services

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"golang.org/x/oauth2"
)

func ListReposInGithub(language string, page int) ([]github.Repository, error) {
	ctx := context.Background()
	staticToken := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: utils.GetGithubAPIKey()})
	tokenClient := oauth2.NewClient(ctx, staticToken)
	client := github.NewClient(tokenClient)

	options := &github.SearchOptions{
		Order: "desc",
		Sort: "stars",
		ListOptions: github.ListOptions{
			Page: 		page,
			PerPage:	 50,
		},
	}

	query := "language:"+language

	result, _, err := client.Search.Repositories(ctx, query, options)

	if err != nil {
		return nil, err
	}

	return result.Repositories, nil
}