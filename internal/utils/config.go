package utils

import (
	"os"
)

func GetGithubAPIKey() string {
	return os.Getenv("GITHUB_API_KEY")
}