package repositories

import (
	"bytes"
	"fmt"
	"house-scanner-backend/internal/db"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/supabase-community/supabase-go"
)

type FileStoreRepository struct {
	supabase *supabase.Client
}

func NewFileStoreRepository() *FileStoreRepository {
	return &FileStoreRepository{supabase: db.GetSupabaseClient()}
}

func (r *FileStoreRepository) UploadFile(fileContent []byte, bucketName string, filePath string) error {
	reader := bytes.NewReader(fileContent)

	// Create request manually to ensure correct Content-Length
	supabaseUrl := os.Getenv("SUPABASE_STORAGE_URL")
	if !strings.HasPrefix(supabaseUrl, "https://") {
		supabaseUrl = "https://" + supabaseUrl
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseUrl, bucketName, filePath), reader)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(fileContent)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SUPABASE_STORAGE_KEY")))

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("upload failed: %v", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

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
