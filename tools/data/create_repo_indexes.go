package tools_data

import (
	"context"

	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRepoIndexes(client *mongo.Client) error {
	db := client.Database(utils.MongoDbDatabase())

	db.Collection("col_repo").Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		mongo.IndexModel{Keys: bson.M{"name": 1}},
		mongo.IndexModel{Keys: bson.M{"email": 1}},
	}, &options.CreateIndexesOptions{})

	return nil
}