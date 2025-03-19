package repositories

import (
	"house-scanner-backend/internal/models"

	"gorm.io/gorm"
)

type FileStoreRepository struct {
	db *gorm.DB
}

func NewFileStoreRepository(db *gorm.DB) *FileStoreRepository {
	return &FileStoreRepository{db: db}
}

func (r *FileStoreRepository) CreateFile(file *models.File) error {
	return r.db.Create(file).Error
}

func (r *FileStoreRepository) GetFile(id string) (*models.File, error) {
	var file models.File
	if err := r.db.Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *FileStoreRepository) DeleteFile(id string) error {
	return r.db.Delete(&models.File{}, id).Error
}

func (r *FileStoreRepository) UploadFile(file *models.File) error {
	return r.db.Create(file).Error
}
