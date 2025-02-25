package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type HouseService struct {
	repo *repositories.HouseRepository
}

func NewHouseService(repo *repositories.HouseRepository) *HouseService {
	return &HouseService{repo: repo}
}

func (s *HouseService) GetHouseByID(id int) (*models.House, error) {
	return s.repo.GetHouseByID(id)
} 