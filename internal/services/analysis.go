package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type AnalysisService struct {
	repo *repositories.AnalysisRepository
}

func NewAnalysisService() *AnalysisService {
	return &AnalysisService{repo: repositories.NewAnalysisRepository()}
}

func (s *AnalysisService) CreateAnalysis(analysis *models.Analysis) error {
	return s.repo.CreateAnalysis(analysis)
}

func (s *AnalysisService) GetAnalysis(id string) (*models.Analysis, error) {
	return s.repo.GetAnalysis(id)
}

func (s *AnalysisService) UpdateAnalysis(id string, analysis *models.Analysis) error {
	return s.repo.UpdateAnalysis(id, analysis)
}

func (s *AnalysisService) DeleteAnalysis(id string) error {
	return s.repo.DeleteAnalysis(id)
}

func (s *AnalysisService) GetAnalyses(name string, phone string) ([]models.Analysis, error) {
	return s.repo.GetAnalyses(name, phone)
}
