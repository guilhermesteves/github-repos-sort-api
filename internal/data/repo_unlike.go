package data

import (
	"context"
	"time"

	"github.com/guilhermesteves/github-repos-sort-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UnlikeRepoById(id string, params map[string]string) (bool, error) {
	client, err := utils.ConnectOnMongo()

	if err != nil {
		return false, err
	}

	db := client.Database(utils.MongoDbDatabase())
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	ip := params["ip"]

	upsert := false

	_, err = db.Collection("col_repo").UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$inc": bson.M{"likes": int64(-1)},
			"$pull": bson.M{"likedBy": ip},
		},
		&options.UpdateOptions{Upsert: &upsert},
	)

	if err != nil {
		return false, err
	}

	return false, nil
}
