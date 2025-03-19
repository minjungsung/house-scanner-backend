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

func (s *FileStoreService) CreateFile(file *models.File, bucketName string, filePath string, data []byte) error {
	return s.repo.CreateFile(file, bucketName, filePath, data)
}

func (s *FileStoreService) GetFile(bucketName string, filePath string) ([]byte, error) {
	return s.repo.GetFile(bucketName, filePath)
}

func (s *FileStoreService) DeleteFile(bucketName string, filePath string) error {
	return s.repo.DeleteFile(bucketName, filePath)
}

func (s *FileStoreService) UploadFile(file *models.File, bucketName string, filePath string, data []byte) error {
	return s.repo.UploadFile(file, bucketName, filePath, data)
}
