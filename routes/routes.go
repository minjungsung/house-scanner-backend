package routes

import (
	"house-scanner-backend/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	post := api.Group("/posts")
	post.Get("/", controllers.GetPosts)
	post.Post("/", controllers.CreatePost)
	post.Get("/:id", controllers.GetPost)
	post.Put("/:id", controllers.UpdatePost)
	post.Delete("/:id", controllers.DeletePost)
	post.Post("/:id/view", controllers.IncreaseView)
	post.Post("/:id/like", controllers.IncreaseLike)
	post.Post("/:id/unlike", controllers.DecreaseLike)

	comment := api.Group("/comments")
	comment.Post("/", controllers.CreateComment)
	comment.Get("/:id", controllers.GetComment)
	comment.Put("/:id", controllers.UpdateComment)
	comment.Delete("/:id", controllers.DeleteComment)

	user := api.Group("/users")
	user.Post("/register", controllers.RegisterUser)
	user.Get("/:email", controllers.GetUserByEmail)
	user.Post("/login", controllers.Login)
	user.Post("/logout", controllers.Logout)
	user.Post("/signup", controllers.SignUp)

	analysis := api.Group("/analyses")
	analysis.Post("/", controllers.CreateAnalysis)
	analysis.Get("/:id", controllers.GetAnalysis)
	analysis.Get("/", controllers.GetAnalyses)
	analysis.Put("/:id", controllers.UpdateAnalysis)
	analysis.Delete("/:id", controllers.DeleteAnalysis)
}
