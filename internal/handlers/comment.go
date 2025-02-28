package handlers

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	CommentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{CommentService: commentService}
}

func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	return nil
}

func GetComment(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid comment ID"})
	}

	comment, err := services.NewCommentService(repositories.NewCommentRepository(db.GetPostgresDB())).GetComment(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Comment not found"})
	}

	return c.JSON(comment)
}

func UpdateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	return nil
}

func DeleteComment(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid comment ID"})
	}

	err = services.NewCommentService(repositories.NewCommentRepository(db.GetPostgresDB())).DeleteComment(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete comment"})
	}

	return nil
}
