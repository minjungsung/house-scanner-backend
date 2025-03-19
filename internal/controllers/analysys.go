package controllers

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/services"
	"strconv"

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

	// Validate required fields
	if analysis.Name == "" || analysis.Phone == "" || analysis.Email == "" || analysis.RequestType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Create analysis in database
	if err := services.NewAnalysisService(repositories.NewAnalysisRepository()).CreateAnalysis(analysis); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create analysis"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Analysis created successfully"})
}

func GetAnalysis(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid analysis ID"})
	}

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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid analysis ID"})
	}

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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid analysis ID"})
	}

	if err := services.NewAnalysisService(repositories.NewAnalysisRepository()).DeleteAnalysis(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete analysis"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Analysis deleted successfully"})
}
