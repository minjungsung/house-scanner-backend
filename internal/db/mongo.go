package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneToMongo(client *mongo.Client, database, collection string, data interface{}) error {
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

func FindOneFromMongo(client *mongo.Client, database, collection string, filter interface{}) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Data retrieved from MongoDB")
	return result, nil
}

func DeleteDataFromMongo(client *mongo.Client, database, collection string, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	log.Println("✅ Data deleted from MongoDB")
	return nil
}

func UpdateDataInMongo(client *mongo.Client, database, collection string, filter, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Println("✅ Data updated in MongoDB")
	return nil
}

func InsertManyToMongo(client *mongo.Client, database, collection string, data []interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	log.Println("✅ Multiple documents inserted into MongoDB")
	return nil
}

func FindManyFromMongo(client *mongo.Client, database, collection string, filter interface{}) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	log.Println("✅ Multiple documents retrieved from MongoDB")
	return results, nil
}

func UpdateManyInMongo(client *mongo.Client, database, collection string, filter, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	log.Println("✅ Multiple documents updated in MongoDB")
	return nil
}

func DeleteManyFromMongo(client *mongo.Client, database, collection string, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := client.Database(database).Collection(collection)
	_, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	log.Println("✅ Multiple documents deleted from MongoDB")
	return nil
}
