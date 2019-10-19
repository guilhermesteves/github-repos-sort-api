package api

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/data"
	"github.com/guilhermesteves/github-repos-sort-api/internal/services"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
)

func ListRepos(lang string, page int) (string, error) {
	lang = validate(lang)

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

func validate(language string) string {
	if language == "" {
		language = "go"
	}

	accepted := []string{
		"javascript", "python", "java", "ruby", "php", "c++", "c#", "go", "c", "swift", "scala", "objective-c",
	}

	for _, l := range accepted {
		if l == language {
			return l
		}
	}
	return "go"
}