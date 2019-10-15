package tools_data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/guilhermesteves/github-repos-sort-api/internal/schema"
	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRepoValidationSchemas(client *mongo.Client) error {
	db := client.Database(utils.MongoDbDatabase())

	jsonSchema := bson.M{}
	jsonFile, _ := schema.ApiDataSchemaRepoJsonBytes()

	err := json.Unmarshal(jsonFile, &jsonSchema)

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	r := db.RunCommand(ctx, bson.D{
		{"collMod", "col_repo"},
		{"validator", bson.M{"$jsonSchema": jsonSchema}},
		{"validationLevel", "moderate"},
	}, &options.RunCmdOptions{})

	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
