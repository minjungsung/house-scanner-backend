package repositories

import (
	"house-scanner-backend/internal/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetComment(id int) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Model(&models.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

func (r *CommentRepository) DeleteComment(id int) error {
	return r.db.Delete(&models.Comment{}, id).Error
}

func (r *CommentRepository) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
