package main

import (
	"house-scanner-backend/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	if err := config.CheckPostgresConnection(cfg.PostgresDSN); err != nil {
		log.Fatalf("PostgreSQL connection error: %v", err)
	}

	if err := config.CheckMongoConnection(cfg.MongoDSN); err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
