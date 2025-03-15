package repositories

import (
	"house-scanner-backend/internal/models"

	"gorm.io/gorm"
)

type AnalysisRepository struct {
	db *gorm.DB
}

func NewAnalysisRepository(db *gorm.DB) *AnalysisRepository {
	return &AnalysisRepository{db: db}
}

func (r *AnalysisRepository) CreateAnalysis(analysis *models.Analysis) error {
	return r.db.Create(analysis).Error
}

func (r *AnalysisRepository) GetAnalysis(id int) (*models.Analysis, error) {
	var analysis models.Analysis
	if err := r.db.Where("id = ?", id).First(&analysis).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *AnalysisRepository) UpdateAnalysis(id int, analysis *models.Analysis) error {
	return r.db.Model(&models.Analysis{}).Where("id = ?", id).Updates(analysis).Error
}

func (r *AnalysisRepository) DeleteAnalysis(id int) error {
	return r.db.Delete(&models.Analysis{}, id).Error
}

func (r *AnalysisRepository) GetAnalyses() ([]models.Analysis, error) {
	var analyses []models.Analysis
	if err := r.db.Find(&analyses).Error; err != nil {
		return nil, err
	}
	return analyses, nil
}
