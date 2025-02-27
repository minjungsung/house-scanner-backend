package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/supabase-community/supabase-go"
)

type Config struct {
	ServerAddress string
	PostgresDSN   string
	MongoDSN      string
	SupabaseURL   string
	SupabaseKey   string
}


func LoadConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		PostgresDSN:   getEnv("POSTGRES_DSN", "postgres://admin:secret@localhost:5432/house_scanner?sslmode=disable"),
		MongoDSN:      getEnv("MONGO_DSN", "mongodb+srv://<username>:<password>@cluster0.mongodb.net/<dbname>?retryWrites=true&w=majority"),
		SupabaseURL:   getEnv("SUPABASE_API_URL", "https://your-supabase-url.supabase.co"),
		SupabaseKey:   getEnv("SUPABASE_KEY", "your-supabase-api-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Reusable MongoDB client
var mongoClient *mongo.Client

func GetMongoClient(dsn string) (*mongo.Client, error) {
	if mongoClient != nil {
		return mongoClient, nil
	}

	clientOptions := options.Client().ApplyURI(dsn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("✅ Successfully connected to MongoDB")
	mongoClient = client
	return mongoClient, nil
}

// Reusable Supabase client

func GetSupabaseClient() *supabase.Client {
	apiUrl := os.Getenv("SUPABASE_API_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")

	if apiUrl == "" || apiKey == "" {
		log.Fatal("Supabase API URL or API Key is not set")
	}

	client, err := supabase.NewClient(apiUrl, apiKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatalf("Cannot initialize Supabase client: %v", err)
	}

	log.Println("✅ Successfully initialized Supabase client")
	return client
}