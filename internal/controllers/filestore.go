package controllers

import (
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FileStoreHandler struct {
	FileStoreService *services.FileStoreService
}

func NewFileStoreHandler(fileStoreService *services.FileStoreService) *FileStoreHandler {
	return &FileStoreHandler{FileStoreService: fileStoreService}
}

func UploadFile(c *fiber.Ctx) error {
	bucketName := c.FormValue("bucket")
	if bucketName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bucket name is required",
		})
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer src.Close()

	// Generate a unique filename
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	// Upload to Supabase Storage
	supabase := db.GetSupabaseClient()
	fileURL := supabase.Storage.GetPublicUrl(bucketName, filename)

	// Upload file using REST API
	_, err = supabase.Storage.UploadFile(bucketName, filename, src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url":         fileURL,
		"filename":    filename,
		"uploaded_at": time.Now(),
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
