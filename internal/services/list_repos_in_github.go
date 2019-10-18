package services

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/guilhermesteves/github-repos-sort-api/internal/model"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"golang.org/x/oauth2"
)

func ListReposInGithub(language string, page int) ([]*model.Repo, error) {
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

	var repos []*model.Repo

	for _, repo := range result.Repositories {
		pushedAt := *repo.PushedAt

		githubRepo := &model.GithubRepo{
			Id: 			*repo.ID,
			FullName: 		*repo.FullName,
			Name:			*repo.Name,
			Description:	*repo.Description,
			Homepage: 		*repo.Homepage,
			DefaultBranch: 	*repo.DefaultBranch,
			PushedAt: 		utils.ToFormattedTime(pushedAt.String()),
			CloneUrl: 		*repo.CloneURL,
			HtmlUrl: 		*repo.HTMLURL,
			Forks: 			int64(*repo.ForksCount),
			Stars: 			int64(*repo.StargazersCount),
		}

		repos = append(repos, &model.Repo{
			Id:			*repo.NodeID,
			GithubRepo: githubRepo,
		})
	}

	return repos, nil
}