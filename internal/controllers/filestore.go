package controllers

import (
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FileStoreHandler struct {
	FileStoreService *services.FileStoreService
}

func NewFileStoreHandler(fileStoreService *services.FileStoreService) *FileStoreHandler {
	return &FileStoreHandler{FileStoreService: fileStoreService}
}

func UploadFile(c *fiber.Ctx) error {
	// Get raw file content from request body
	fileContent := c.Body()

	// Get filename from header
	fileName := c.Get("X-File-Name")
	if fileName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "X-File-Name header is required",
		})
	}

	// Upload to Supabase Storage
	err := services.NewFileStoreService(repositories.NewFileStoreRepository()).UploadFile(fileContent, "documents", fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to upload file: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "File uploaded successfully",
		"fileName": fileName,
		"size":     len(fileContent),
		"path":     fmt.Sprintf("documents/%s", fileName),
	})
}

func GetFile(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	filePath := c.Params("path")

	file, err := services.NewFileStoreService(repositories.NewFileStoreRepository()).GetFile(bucketName, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"file": file})
}

func DeleteFile(c *fiber.Ctx) error {
	bucketName := c.Params("bucket")
	filePath := c.Params("path")

	err := services.NewFileStoreService(repositories.NewFileStoreRepository()).DeleteFile(bucketName, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully"})
}
