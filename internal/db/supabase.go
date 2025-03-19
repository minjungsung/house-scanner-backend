package db

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

var supabaseClient *supabase.Client

func GetSupabaseClient() *supabase.Client {
	if supabaseClient == nil {
		baseURL := os.Getenv("SUPABASE_API_URL")
		apiKey := os.Getenv("SUPABASE_API_KEY")

		// Storage API URL 형식으로 변경
		storageURL := baseURL + "/storage/v1"
		fmt.Println("SUPABASE_URL", storageURL)
		fmt.Println("SUPABASE_KEY", apiKey)

		client, err := supabase.NewClient(storageURL, apiKey, nil)
		if err != nil {
			panic(err)
		}
		supabaseClient = client
	}
	return supabaseClient
}
