package main

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	db.GetPostgresDB() // PostgreSQL 연결
	db.GetMongoDB()    // MongoDB 연결

	app := fiber.New(fiber.Config{
		AppName: "House Scanner API",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	routes.SetupRoutes(app)

	// Create Socket.IO server with custom transport options
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})
	defer server.Close()

	// Socket.IO 이벤트 처리
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Printf("Client connected: %s", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("Client disconnected: %s, reason: %s", s.ID(), reason)
	})

	// 동적 이벤트 구독을 위한 핸들러
	server.OnEvent("/", "subscribe_analysis", func(s socketio.Conn, analysisID string) {
		log.Printf("Client %s subscribed to analysis: %s", s.ID(), analysisID)
	})

	server.OnEvent("/", "unsubscribe_analysis", func(s socketio.Conn, analysisID string) {
		log.Printf("Client %s unsubscribed from analysis: %s", s.ID(), analysisID)
	})

	// Socket.IO 서버 시작 (별도 포트)
	go func() {
		http.Handle("/socket.io/", server)
		log.Printf("Socket.IO server starting on port 8081...")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Error starting Socket.IO server: %v", err)
		}
	}()

	// API 서버 시작
	log.Printf("API server starting on port 8080...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting API server: %v", err)
	}
}
