package models

import "time"

type File struct {
	Filename         string    `gorm:"not null"`
	Filepath         string    `gorm:"not null"`
	FileData         []byte    `gorm:"not null"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime"`
	UpdatedTimestamp time.Time `gorm:"autoUpdateTime"`
}
