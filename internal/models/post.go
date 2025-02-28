package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Views     int       `gorm:"default:0"`
	Likes     int       `gorm:"default:0"`
}
