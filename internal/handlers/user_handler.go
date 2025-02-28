package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUser(mongoClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request to register user")

		// Log the raw request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		log.Printf("Raw request body: %s\n", body)

		// Reset the request body for decoding
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Decoded user data: %+v\n", user)

		// Validate input
		if user.Name == "" || user.Phone == "" || user.Email == "" || user.Address == "" || user.Message == "" || len(user.ReferralSource) == 0 {
			log.Println("Missing required fields")
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Insert into MongoDB
		log.Println("Inserting user into MongoDB")
		if err := db.InsertOneToMongo(mongoClient, "house_scanner", "users", user); err != nil {
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
