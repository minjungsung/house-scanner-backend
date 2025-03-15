package models

import "time"

type Analysis struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Phone       string    `gorm:"not null"`
	Email       string    `gorm:"not null"`
	RequestType string    `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type AnalysisStatus string

const (
	Pending   AnalysisStatus = "pending"
	Approved  AnalysisStatus = "approved"
	Rejected  AnalysisStatus = "rejected"
	Completed AnalysisStatus = "completed"
	Failed    AnalysisStatus = "failed"
	Canceled  AnalysisStatus = "canceled"
	Processing AnalysisStatus = "processing"
)
