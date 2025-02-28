package services

import (
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(email string) error {
	return s.repo.DeleteUser(email)
}
