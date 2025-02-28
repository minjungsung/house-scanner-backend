package main

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.GetPostgresDB() // PostgreSQL 연결
	db.GetMongoDB()    // MongoDB 연결

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
