package routes

import (
	"house-scanner-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/register", handlers.RegisterUser)
	user.Get("/:email", handlers.GetUserByEmail)

	post := api.Group("/posts")
	post.Post("/", handlers.CreatePost)
	post.Get("/:id", handlers.GetPost)
	post.Put("/:id", handlers.UpdatePost)
	post.Delete("/:id", handlers.DeletePost)

	comment := api.Group("/comments")
	comment.Post("/", handlers.CreateComment)
	comment.Get("/:id", handlers.GetComment)
	comment.Put("/:id", handlers.UpdateComment)
	comment.Delete("/:id", handlers.DeleteComment)
}
