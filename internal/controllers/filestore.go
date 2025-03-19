package controllers

import (
	"fmt"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type FileStoreHandler struct {
	FileStoreService *services.FileStoreService
}

func NewFileStoreHandler(fileStoreService *services.FileStoreService) *FileStoreHandler {
	return &FileStoreHandler{FileStoreService: fileStoreService}
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse file"})
	}

	filename := file.Filename
	filepath := fmt.Sprintf("uploads/%s", filename)

	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File uploaded successfully"})
}

func GetFile(c *fiber.Ctx) error {
	id := c.Params("id")

	file, err := services.NewFileStoreService(repositories.NewFileStoreRepository(db.GetPostgresDB())).GetFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"file": file})
}

func DeleteFile(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.NewFileStoreService(repositories.NewFileStoreRepository(db.GetPostgresDB())).DeleteFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully"})
}
