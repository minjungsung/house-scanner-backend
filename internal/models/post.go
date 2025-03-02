package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Views     int       `gorm:"default:0"`
	Likes     int       `gorm:"default:0"`
}

// Custom JSON response
type PostResponse struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"` // Maps to UpdatedAt
	Views    int       `json:"views"`
	Likes    int       `json:"likes"`
}

// Convert Post to PostResponse
func (p *Post) ToResponse() *PostResponse {
	return &PostResponse{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
		Date:    p.UpdatedAt, // Use UpdatedAt for Date
		Views:   p.Views,
		Likes:   p.Likes,
	}
}
