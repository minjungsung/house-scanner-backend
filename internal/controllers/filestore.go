package controllers

import (
	"house-scanner-backend/internal/services"

	"encoding/base64"
	"fmt"
	"mime"
	"path"
	"strings"

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
	err := services.NewFileStoreService().UploadFile(fileContent, "documents", fileName)
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
	bucketName := "documents"
	fileId := c.Params("id")

	if bucketName == "" || fileId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bucket name and file id are required"})
	}

	// Get file from storage
	fileStore := services.NewFileStoreService()
	fileContent, err := fileStore.GetFile(bucketName, fileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to get file: %v", err),
		})
	}

	// Get MIME type from file extension
	ext := strings.ToLower(path.Ext(fileId))
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		// If extension is not recognized, try to detect from content
		// You might want to use a more sophisticated content detection library here
		mimeType = "application/octet-stream"
	}

	// Base64로 인코딩
	base64Content := base64.StdEncoding.EncodeToString(fileContent)

	return c.JSON(fiber.Map{
		"file":     base64Content,
		"filename": fileId,
		"mimetype": mimeType,
	})
}

func DeleteFile(c *fiber.Ctx) error {
	bucketName := "documents"
	fileId := c.Params("id")

	err := services.NewFileStoreService().DeleteFile(bucketName, fileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete file"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "File deleted successfully"})
}
