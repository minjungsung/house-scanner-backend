package routes

import (
	"house-scanner-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/register", handlers.RegisterUserHandler)
	user.Get("/:email", handlers.GetUserByEmailHandler)
}
