package controllers

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"

	"io"

	"github.com/gofiber/fiber/v2"
)

type AnalysisHandler struct {
	AnalysisService *services.AnalysisService
}

func NewAnalysisHandler(analysisService *services.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{AnalysisService: analysisService}
}

func CreateAnalysis(c *fiber.Ctx) error {
	// Create new analysis instance
	analysis := new(models.Analysis)
	analysis.Name = c.FormValue("name")
	analysis.Phone = c.FormValue("phone")
	analysis.Email = c.FormValue("email")
	analysis.RequestType = c.FormValue("requestType")

	// Handle file upload
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload failed"})
	}

	// Open and read the file
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file"})
	}


	// Validate required fields
	if analysis.Name == "" || analysis.Phone == "" || analysis.Email == "" || analysis.RequestType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Upload file to Supabase Storage
	if err := services.NewFileStoreService(repositories.NewFileStoreRepository()).UploadFile(fileContent, "documents", file.Filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file"})
	}

	// Create analysis in database
	if err := services.NewAnalysisService(repositories.NewAnalysisRepository()).CreateAnalysis(analysis); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create analysis"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Analysis created successfully"})
}

func GetAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	analysis, err := services.NewAnalysisService(repositories.NewAnalysisRepository()).GetAnalysis(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Analysis not found"})
	}

	return c.Status(fiber.StatusOK).JSON(analysis)
}

func GetAnalyses(c *fiber.Ctx) error {
	analyses, err := services.NewAnalysisService(repositories.NewAnalysisRepository()).GetAnalyses()
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

	if err := services.NewAnalysisService(repositories.NewAnalysisRepository()).UpdateAnalysis(id, analysis); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis updated successfully"})
}

func DeleteAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id := idStr

	if err := services.NewAnalysisService(repositories.NewAnalysisRepository()).DeleteAnalysis(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis deleted successfully"})
}
