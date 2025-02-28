package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.CreatePost(post)
}

func (s *PostService) GetPost(id int) (*models.Post, error) {
	return s.repo.GetPost(id)
}

func (s *PostService) UpdatePost(post *models.Post) error {
	return s.repo.UpdatePost(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) GetPostByID(id int) (*models.Post, error) {
	return s.repo.GetPostByID(id)
}
