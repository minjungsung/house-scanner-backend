package db

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

// Supabase 데이터 삽입 함수
func InsertDataToSupabase(client *supabase.Client, table string, data interface{}) error {
	// Assuming the correct method chain for inserting data
	_, _, err := client.From(table).Insert(data, false, "", "", "").Execute()
	if err != nil {
		return err
	}

	log.Println("✅ Data inserted into Supabase")
	return nil
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
