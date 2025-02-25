package repositories

import (
	"database/sql"
	"house-scanner-backend/internal/models"
)

type HouseRepository struct {
	db *sql.DB
}

func NewHouseRepository(db *sql.DB) *HouseRepository {
	return &HouseRepository{db: db}
}

func (r *HouseRepository) GetHouseByID(id int) (*models.House, error) {
	// Implement database logic to get a house by ID
	return nil, nil
} 