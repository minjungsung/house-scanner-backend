package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	onceMongo   sync.Once
)

func GetMongoDB() *mongo.Client {
	onceMongo.Do(func() {
		uri := os.Getenv("MONGO_DSN")
		if uri == "" {
			log.Fatal("❌ MongoDB URI is not set")
		}

		clientOptions := options.Client().ApplyURI(uri)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		MongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
		}

		err = MongoClient.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("❌ MongoDB connection test failed: %v", err)
		}

		log.Println("✅ Successfully connected to MongoDB")
	})

	return MongoClient
}
