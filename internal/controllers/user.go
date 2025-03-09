package controllers

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"

	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewErrorResponse(models.ERROR_INVALID_INPUT))
	}

	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))
	if err := userService.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewErrorResponse(models.ERROR_INTERNAL_SERVER))
	}

	log.Printf("User registered/updated: %s\n", user.Email)
	return c.Status(fiber.StatusOK).JSON(models.NewSuccessResponse(user))
}

func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))
	user, err := userService.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.NewErrorResponse(models.ERROR_USER_NOT_FOUND))
	}

	return c.JSON(models.NewSuccessResponse(user))
}

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// 🟢 회원가입 (Signup)
func (c *UserController) Signup(ctx *fiber.Ctx) error {
	var input struct {
		Name          string `json:"name"`
		Email         string `json:"email"`
		Password      string `json:"password"`
		Address       string `json:"address"`
		AddressDetail string `json:"address_detail"`
		Birthday      string `json:"birthday"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user := &models.User{
		Email:          input.Email,
		HashedPassword: input.Password,
		Name:           input.Name,
		Address:        input.Address,
		AddressDetail:  input.AddressDetail,
		Birthday:       input.Birthday,
	}

	err := c.UserService.SignupUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User signed up successfully"})
}

// 🟢 로그인 (Login)
func (c *UserController) Login(ctx *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := c.UserService.LoginUser(input.Email, input.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}

func SignUp(c *fiber.Ctx) error {
	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	err := userService.SignupUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Signed up successfully"})
}

func Login(c *fiber.Ctx) error {
	userService := services.NewUserService(repositories.NewUserRepository(db.GetPostgresDB()))

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	token, err := userService.LoginUser(user.Email, user.HashedPassword)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
