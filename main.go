package main

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.GetPostgresDB() // PostgreSQL 연결
	db.GetMongoDB()    // MongoDB 연결

	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // Allow all origins
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
