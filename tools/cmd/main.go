package main

import (
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"log"

	"github.com/guilhermesteves/github-repos-sort-api/tools/data"
)

func main() {
	// FIXME: ensure collection via each schema
	db, err := utils.ConnectOnMongo()

	if err != nil {
		panic(err)
	}

	tools_data.CreateRepoIndexes(db)
	tools_data.CreateRepoValidationSchemas(db)

	utils.DisconnectFromMongo(db)
	log.Println("db is updated")
}
