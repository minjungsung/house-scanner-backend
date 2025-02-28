package repositories

import (
	"house-scanner-backend/internal/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetPost(id int) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(post).Error
}

func (r *PostRepository) DeletePost(id int) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
