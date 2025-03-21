package services

import (
	"house-scanner-backend/internal/repositories"
)

type FileStoreService struct {
	repo *repositories.FileStoreRepository
}

func NewFileStoreService() *FileStoreService {
	repo := repositories.NewFileStoreRepository()
	return &FileStoreService{repo: repo}
}

func (s *FileStoreService) UploadFile(fileContent []byte, bucketName string, filePath string) error {
	return s.repo.UploadFile(fileContent, bucketName, filePath)
}

func (s *FileStoreService) GetFile(bucketName string, filePath string) ([]byte, error) {
	return s.repo.GetFile(bucketName, filePath)
}

func (s *FileStoreService) DeleteFile(bucketName string, filePath string) error {
	return s.repo.DeleteFile(bucketName, filePath)
}
