package db

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

var supabaseClient *supabase.Client

func GetSupabaseClient() *supabase.Client {
	if supabaseClient == nil {
		supabaseURL := os.Getenv("SUPABASE_STORAGE_URL")
		supabaseKey := os.Getenv("SUPABASE_STORAGE_KEY")

		client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
		if err != nil {
			panic(err)
		}
		supabaseClient = client
	}
	return supabaseClient
}
