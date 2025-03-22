package db

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

var supabaseClient *supabase.Client

func GetSupabaseClient() *supabase.Client {
	if supabaseClient == nil {
		baseURL := os.Getenv("SUPABASE_URL")
		apiKey := os.Getenv("SUPABASE_KEY")

		// Storage API URL 형식으로 변경
		storageURL := baseURL + "/storage/v1"
		client, err := supabase.NewClient(storageURL, apiKey, nil)
		if err != nil {
			panic(err)
		}
		supabaseClient = client
	}
	return supabaseClient
}
