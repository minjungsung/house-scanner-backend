package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type CommentService struct {
	repo *repositories.CommentRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetComment(id int) (*models.Comment, error) {
	return s.repo.GetComment(id)
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
	return s.repo.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(id int) error {
	return s.repo.DeleteComment(id)
}

func (s *CommentService) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return s.repo.GetCommentsByPostID(postID)
}
