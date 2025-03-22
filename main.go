package main

import (
	"context"
	"encoding/json"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/routes"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// WebSocket 클라이언트 구조체
type Client struct {
	ID      string
	Conn    *websocket.Conn
	Send    chan []byte
	EventID string
}

// Hub는 모든 활성 클라이언트와 브로드캐스트 메시지를 관리합니다
type Hub struct {
	clients    map[string]*Client
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.RWMutex
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.ID] = client
			log.Printf("[WebSocket] Client registered. Total clients: %d", len(h.clients))
			h.mutex.Unlock()
			log.Printf("[WebSocket] Client connected: %s", client.ID)

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client.ID]; ok {
				delete(h.clients, client.ID)
				log.Printf("[WebSocket] Client unregistered. Total clients: %d", len(h.clients))
			}
			h.mutex.Unlock()
			close(client.Send)
			log.Printf("[WebSocket] Client disconnected: %s", client.ID)

		case message := <-h.broadcast:
			h.mutex.RLock()
			log.Printf("[WebSocket] Broadcasting to %d clients", len(h.clients))
			for _, client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client.ID)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[WebSocket] Error reading message: %v", err)
			}
			break
		}

		var msg struct {
			Type    string `json:"type"`
			EventID string `json:"eventId"`
		}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("[WebSocket] Error parsing message: %v", err)
			continue
		}

		switch msg.Type {
		case "subscribe":
			c.EventID = msg.EventID
			response := map[string]interface{}{
				"type":      "subscribed",
				"eventId":   msg.EventID,
				"clientId":  c.ID,
				"timestamp": time.Now().Unix(),
			}
			if err := c.Conn.WriteJSON(response); err != nil {
				log.Printf("[WebSocket] Error sending subscription confirmation: %v", err)
			}

		case "unsubscribe":
			c.EventID = ""
			response := map[string]interface{}{
				"type":      "unsubscribed",
				"eventId":   msg.EventID,
				"clientId":  c.ID,
				"timestamp": time.Now().Unix(),
			}
			if err := c.Conn.WriteJSON(response); err != nil {
				log.Printf("[WebSocket] Error sending unsubscription confirmation: %v", err)
			}
		}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func main() {
	// 종료 시그널 처리
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	db.GetPostgresDB()         // PostgreSQL 연결
	mongoDB := db.GetMongoDB() // MongoDB 연결

	app := fiber.New(fiber.Config{
		AppName: "House Scanner API",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,http://127.0.0.1:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
		MaxAge:           43200, // 12 hours in seconds
	}))

	// WebSocket Hub 생성
	hub := newHub()
	go hub.run()

	// WebSocket 엔드포인트 설정
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		log.Printf("[WebSocket] New connection attempt from: %s", c.RemoteAddr().String())

		// Set read deadline
		c.SetReadLimit(512 * 1024) // 512KB
		c.SetReadDeadline(time.Now().Add(60 * time.Second))
		c.SetPongHandler(func(string) error {
			c.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})

		client := &Client{
			ID:   c.RemoteAddr().String(),
			Conn: c,
			Send: make(chan []byte, 256),
		}

		log.Printf("[WebSocket] Sending client to registration channel")
		hub.register <- client

		// 연결 성공 메시지 전송
		response := map[string]interface{}{
			"type":      "connected",
			"clientId":  client.ID,
			"timestamp": time.Now().Unix(),
		}
		if err := c.WriteJSON(response); err != nil {
			log.Printf("[WebSocket] Error sending connection success: %v", err)
		}

		// Start ping ticker
		go func() {
			ticker := time.NewTicker(30 * time.Second)
			defer func() {
				ticker.Stop()
				c.Close()
			}()

			for {
				select {
				case <-ticker.C:
					if c.Conn == nil {
						return
					}
					if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
						log.Printf("[WebSocket] Error sending ping: %v", err)
						return
					}
				default:
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()

		go client.writePump()
		go client.readPump(hub)
	}))

	routes.SetupRoutes(app)

	// MongoDB Change Stream 처리
	go func() {
		collection := mongoDB.Database("house_scanner").Collection("analysis")
		pipeline := mongo.Pipeline{
			{{"$match", bson.D{{"operationType", bson.D{{"$in", []string{"insert", "update"}}}}}}},
		}

		changeStream, err := collection.Watch(context.Background(), pipeline)
		if err != nil {
			log.Fatalf("Error creating change stream: %v", err)
		}
		defer changeStream.Close(context.Background())

		log.Println("[MongoDB] Change Stream started...")
		for changeStream.Next(context.Background()) {
			var changeEvent struct {
				DocumentKey struct {
					ID string `bson:"_id"`
				} `bson:"documentKey"`
				FullDocument  bson.M `bson:"fullDocument"`
				OperationType string `bson:"operationType"`
			}
			if err := changeStream.Decode(&changeEvent); err != nil {
				log.Printf("[MongoDB] Error decoding change event: %v", err)
				continue
			}

			eventID := changeEvent.DocumentKey.ID
			log.Printf("[MongoDB] Change detected for event: %s, operation: %s", eventID, changeEvent.OperationType)

			// 변경사항을 구독 중인 클라이언트에게 전송
			updateData := map[string]interface{}{
				"type":      "update",
				"eventId":   eventID,
				"data":      changeEvent.FullDocument,
				"operation": changeEvent.OperationType,
				"timestamp": time.Now().Unix(),
			}

			message, err := json.Marshal(updateData)
			if err != nil {
				log.Printf("[WebSocket] Error marshaling update data: %v", err)
				continue
			}

			hub.mutex.RLock()
			for _, client := range hub.clients {
				if client.EventID == eventID {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(hub.clients, client.ID)
					}
				}
			}
			hub.mutex.RUnlock()
		}
	}()

	// API 서버 시작
	go func() {
		log.Printf("API server starting on port 8080...")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Error starting API server: %v", err)
		}
	}()

	// 종료 처리
	<-quit
	log.Println("Shutting down server...")

	// Fiber 서버 종료
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down fiber server: %v", err)
	}

	// MongoDB 연결 종료
	if err := mongoDB.Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}

	log.Println("Server gracefully stopped")
	os.Exit(0)
}
