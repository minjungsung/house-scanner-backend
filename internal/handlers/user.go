package handlers

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"

	"github.com/gofiber/fiber/v2"
	"log"
)

func RegisterUserHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))
	if err := userService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	log.Println("Received user data: %+v\n", user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func GetUserByEmailHandler(c *fiber.Ctx) error {
	email := c.Params("email")

	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))
	user, err := userService.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
