package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

func init() {
	if os.Getenv("ENV") == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
