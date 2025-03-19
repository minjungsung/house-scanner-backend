package models

import "time"

type Analysis struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	Name             string `gorm:"not null"`
	Phone            string `gorm:"not null"`
	Email            string `gorm:"not null"`
	RequestType      string `gorm:"not null"`
	File             File
	Title            string
	Content          string
	CreatedTimestamp time.Time `gorm:"autoCreateTime"`
	UpdatedTimestamp time.Time `gorm:"autoUpdateTime"`
}

type AnalysisStatus string

const (
	Pending    AnalysisStatus = "pending"
	Approved   AnalysisStatus = "approved"
	Rejected   AnalysisStatus = "rejected"
	Completed  AnalysisStatus = "completed"
	Failed     AnalysisStatus = "failed"
	Canceled   AnalysisStatus = "canceled"
	Processing AnalysisStatus = "processing"
)
