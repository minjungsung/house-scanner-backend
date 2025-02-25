package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectMongo(dsn string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB")
	return client, nil
} 