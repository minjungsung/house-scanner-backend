package main

import (
	"log"
	"net/http"

	"house-scanner-backend/config"

	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from Minjung backend!"}`))
}

func main() {
	// 환경 설정 로드
	cfg := config.LoadConfig()

	// PostgreSQL 연결 확인
	if err := config.CheckPostgresConnection(cfg.PostgresDSN); err != nil {
		log.Fatalf("PostgreSQL connection error: %v", err)
	}

	// MongoDB 연결 확인
	if err := config.CheckMongoConnection(cfg.MongoDSN); err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	// HTTP 요청 처리
	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", handler)

	// CORS 설정 추가
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // React 프론트엔드 주소 허용
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	// 서버 시작
	log.Printf("Starting server on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, corsHandler))
}
