package main

import (
	"log"
	"net/http"

	"context"
	"house-scanner-backend/config"
	"house-scanner-backend/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/supabase-community/supabase-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// ✅ 환경 설정 로드
	cfg := config.LoadConfig()

	// ✅ Supabase 클라이언트 초기화
	supabaseClient := config.GetSupabaseClient()
	if supabaseClient == nil {
		log.Fatalf("Supabase client initialization error")
	}

	// ✅ MongoDB 연결
	mongoClient, err := config.GetMongoClient(cfg.MongoDSN)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	// ✅ HTTP 요청 처리
	router := mux.NewRouter()

	// Register routes
	registerRoutes(router, supabaseClient, mongoClient)

	// ✅ CORS 설정 추가
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	// ✅ 서버 시작
	log.Printf("🚀 Starting server on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, corsHandler))
}

func registerRoutes(router *mux.Router, supabaseClient *supabase.Client, mongoClient *mongo.Client) {
	router.HandleFunc("/api/register-user", handlers.RegisterUser(mongoClient)).Methods("POST")
	// Add more routes here as needed
}
