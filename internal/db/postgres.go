package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPostgresDB() *gorm.DB {
	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DSN is not set")
	}

	// Initialize GORM with pgx
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	// Example query to test connection
	var version string
	if err := db.Raw("SELECT version()").Scan(&version).Error; err != nil {
		log.Fatalf("❌ Query failed: %v", err)
	}

	log.Println("✅ Successfully connected to Supabase (PostgreSQL)")

	return db
}
