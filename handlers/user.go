package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"house-scanner-backend/config"

	"github.com/supabase-community/supabase-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at,omitempty"` // Optional if auto-generated
}

func RegisterUser(supabaseClient *supabase.Client, mongoClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request to register user")

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Decoded user data: %+v\n", user)

		// Validate input
		if user.FirstName == "" || user.LastName == "" || user.Email == "" {
			log.Println("Missing required fields")
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Insert into Supabase
		log.Println("Inserting user into Supabase")
		if err := config.InsertDataToSupabase(supabaseClient, "users", user); err != nil {
			log.Printf("Supabase error: %v\n", err)
			http.Error(w, fmt.Sprintf("Supabase error: %v", err), http.StatusInternalServerError)
			return
		}
		log.Println("Successfully inserted user into Supabase")

		// Insert into MongoDB
		log.Println("Inserting user into MongoDB")
		if err := config.InsertDataToMongo(mongoClient, "house_scanner", "users", user); err != nil {
			log.Printf("MongoDB error: %v\n", err)
			http.Error(w, fmt.Sprintf("MongoDB error: %v", err), http.StatusInternalServerError)
			return
		}
		log.Println("Successfully inserted user into MongoDB")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "User registered successfully"}`))
		log.Println("User registration completed successfully")
	}
}
