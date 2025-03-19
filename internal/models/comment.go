// models/comment.go
package models

import "time"

type Comment struct {
	ID               int       `gorm:"primaryKey"`
	PostID           int       `gorm:"not null"`
	Author           string    `gorm:"not null"`
	Content          string    `gorm:"not null"`
	CreatedTimestamp time.Time `gorm:"autoCreateTime"`
}
