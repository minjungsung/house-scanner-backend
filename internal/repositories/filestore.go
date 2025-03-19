package repositories

import (
	"bytes"
	"house-scanner-backend/internal/db"
	"house-scanner-backend/internal/models"

	"github.com/supabase-community/supabase-go"
)

type FileStoreRepository struct {
	supabase *supabase.Client
}

func NewFileStoreRepository() *FileStoreRepository {
	return &FileStoreRepository{supabase: db.GetSupabaseClient()}
}

func (r *FileStoreRepository) CreateFile(file *models.File, bucketName string, filePath string, data []byte) error {
	_ , err := r.supabase.Storage.UploadFile(bucketName, filePath, bytes.NewReader(data))
	return err
}

func (r *FileStoreRepository) GetFile(bucketName string, filePath string) ([]byte, error) {
	data, err := r.supabase.Storage.DownloadFile(bucketName, filePath)
	return data, err
}

func (r *FileStoreRepository) DeleteFile(bucketName string, filePath string) error {
	_, err := r.supabase.Storage.RemoveFile(bucketName, []string{filePath})
	return err
}

func (r *FileStoreRepository) UploadFile(file *models.File, bucketName string, filePath string, data []byte) error {
	return r.CreateFile(file, bucketName, filePath, data)
}
