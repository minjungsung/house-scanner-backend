package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertDataToMongo(client *mongo.Client, database, collection string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	log.Println("✅ Data inserted into MongoDB")
	return nil
}