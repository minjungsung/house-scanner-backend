package repositories

import (
	"bytes"
	"fmt"
	"house-scanner-backend/internal/db"

	"github.com/supabase-community/supabase-go"
)

type FileStoreRepository struct {
	supabase *supabase.Client
}

func NewFileStoreRepository() *FileStoreRepository {
	return &FileStoreRepository{supabase: db.GetSupabaseClient()}
}

func (r *FileStoreRepository) UploadFile(fileContent []byte, bucketName string, filePath string) error {
	reader := bytes.NewBuffer(fileContent)

	response, err := r.supabase.Storage.UploadFile(bucketName, filePath, reader)
	if err != nil {
		fmt.Printf("Upload error: %v\n", err)
		return err
	}
	fmt.Printf("Upload successful. Response: %+v\n", response)
	return nil
}

func (r *FileStoreRepository) GetFile(bucketName string, filePath string) ([]byte, error) {
	data, err := r.supabase.Storage.DownloadFile(bucketName, filePath)
	return data, err
}

func (r *FileStoreRepository) DeleteFile(bucketName string, filePath string) error {
	_, err := r.supabase.Storage.RemoveFile(bucketName, []string{filePath})
	return err
}
