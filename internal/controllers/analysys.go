package controllers

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/services"

	"fmt"
	"io"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AnalysisHandler struct {
	AnalysisService *services.AnalysisService
}

func NewAnalysisHandler(analysisService *services.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{AnalysisService: analysisService}
}

func CreateAnalysis(c *fiber.Ctx) error {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload is required"})
	}

	// Open and read the file
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer src.Close()

	// Read file content
	fileContent, err := io.ReadAll(src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	// Upload file directly using FileStoreService
	fileStore := services.NewFileStoreService()
	err = fileStore.UploadFile(fileContent, "documents", file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to upload file: %v", err),
		})
	}

	// Create new analysis instance
	analysis := new(models.Analysis)
	analysis.Name = c.FormValue("name")
	analysis.Phone = c.FormValue("phone")
	analysis.Email = c.FormValue("email")
	analysis.RequestType = c.FormValue("requestType")
	analysis.File = &models.File{
		Name: file.Filename,
		Size: file.Size,
		Path: fmt.Sprintf("documents/%s", file.Filename),
		Type: file.Header.Get("Content-Type"),
	}

	// Validate required fields
	if analysis.Name == "" || analysis.Phone == "" || analysis.Email == "" || analysis.RequestType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Create analysis in database
	if err := services.NewAnalysisService().CreateAnalysis(analysis, file.Filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create analysis"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Analysis created successfully",
		"analysis": analysis,
	})
}

func GetAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	analysis, err := services.NewAnalysisService().GetAnalysis(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Analysis not found"})
	}

	return c.Status(fiber.StatusOK).JSON(analysis)
}

func GetAnalyses(c *fiber.Ctx) error {
	// should have name and phonenumber
body := new(models.Analysis)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	name := body.Name
	phone := body.Phone

	if name == "" && phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and phone are required"})
	}

	analyses, err := services.NewAnalysisService().GetAnalyses(name, phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve analyses"})
	}

	return c.Status(fiber.StatusOK).JSON(analyses)
}

func UpdateAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	analysis := new(models.Analysis)
	if err := c.BodyParser(analysis); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := services.NewAnalysisService().UpdateAnalysis(id, analysis); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis updated successfully"})
}

func DeleteAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	if err := services.NewAnalysisService().DeleteAnalysis(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis deleted successfully"})
}

func UploadAnalysisFile(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	analysis, err := services.NewAnalysisService().GetAnalysis(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Analysis not found"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload is required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	fileStore := services.NewFileStoreService()
	err = fileStore.UploadFile(fileContent, "documents", file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to upload file: %v", err),
		})
	}

	analysis.AnalysisFileId = file.Filename
	analysis.UpdatedTimestamp = time.Now()
	analysis.Status = models.Completed

	if err := services.NewAnalysisService().UpdateAnalysis(id, analysis); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis file uploaded successfully"})

}
