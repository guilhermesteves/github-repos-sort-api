package api

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/data"
	"github.com/guilhermesteves/github-repos-sort-api/internal/services"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
)

func ListRepos(lang string, page int) (string, error) {
	repos, err := services.ListReposInGithub(lang, page)

	if err != nil {
		return "", err
	}

	_, err = data.SaveAll(repos)

	if err != nil {
		return "", err
	}

	return utils.ApiResult(repos), nil
}