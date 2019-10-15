package utils

import (
	"os"
)

func MongoDbDsn() string {
	return os.Getenv("MONGODB_DSN")
}
func MongoDbDatabase() string {
	return os.Getenv("MONGODB_DATABASE")
}

func GetGithubAPIKey() string {
	return os.Getenv("GITHUB_API_KEY")
}