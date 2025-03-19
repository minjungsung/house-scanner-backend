package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type FileStoreService struct {
	repo *repositories.FileStoreRepository
}

func NewFileStoreService(repo *repositories.FileStoreRepository) *FileStoreService {
	return &FileStoreService{repo: repo}
}

func (s *FileStoreService) CreateFile(file *models.File) error {
	return s.repo.CreateFile(file)
}

func (s *FileStoreService) GetFile(id string) (*models.File, error) {
	return s.repo.GetFile(id)
}

func (s *FileStoreService) DeleteFile(id string) error {
	return s.repo.DeleteFile(id)
}

func (s *FileStoreService) UploadFile(file *models.File) error {
	return s.repo.UploadFile(file)
}
