package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoTries = 0

func ConnectOnMongo() (*mongo.Client, error) {
	opt := &options.ClientOptions{}
	opt = opt.ApplyURI(MongoDbDsn())
	opt = opt.SetMaxPoolSize(3)
	client, err := mongo.NewClient(opt)
	Check(err)

	err = client.Connect(context.TODO())

	if err != nil {
		time.Sleep(time.Second * 1)
		log.Println(err.Error())
		mongoTries++
		if mongoTries > 10 {
			return nil, err
		}
		log.Println("Trying again...")
		return ConnectOnMongo()
	} else {
		return client, nil
	}
}

func DisconnectFromMongo(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())

	if err != nil {
		time.Sleep(time.Second * 1)
		log.Println(err.Error())
		mongoTries++
		if mongoTries > 10 {
			return err
		}
		log.Println("Trying again...")
		return DisconnectFromMongo(client)
	}

	return nil
}

func CursorToArray(ctx context.Context, cursor *mongo.Cursor) bson.A {
	var result = bson.A{}

	for cursor.Next(ctx) {
		elem := &bson.M{}
		if err := cursor.Decode(elem); err != nil {
			log.Println(err)
		} else {
			result = append(result, *elem)
		}
	}

	return result
}
