package api

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/data"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
)

func UnlikeRepo(ip string, repoId string) (string, error) {
	client, err := utils.ConnectOnMongo()

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	defer utils.DisconnectFromMongo(client)

	repo, err := data.UnlikeRepoById(repoId, map[string]string{ "ip": ip })

	if err != nil {
		return "", err
	}

	return utils.ApiResult(repo), nil
}