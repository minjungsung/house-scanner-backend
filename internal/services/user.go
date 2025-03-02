package services

import (
	"errors"
	"house-scanner-backend/internal/models"
	"house-scanner-backend/internal/repositories"
	"house-scanner-backend/internal/utils"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// 🟢 회원가입 (SignupUser)
func (s *UserService) SignupUser(user *models.User) error {
	// ✅ 이메일 중복 체크
	existingUser, _ := s.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already in use")
	}

	// ✅ 비밀번호 해싱 후 저장
	hashedPassword, err := utils.HashPassword(user.HashedPassword)
	if err != nil {
		return err
	}
	user.HashedPassword = hashedPassword

	return s.repo.CreateUser(user)
}

// 🟢 로그인 (JWT 발급)
func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.HashedPassword) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 🟢 유저 정보 가져오기
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

// 🟢 유저 정보 업데이트 (비밀번호 변경 가능)
func (s *UserService) UpdateUser(user *models.User) error {
	if user.HashedPassword != "" {
		hashedPassword, err := utils.HashPassword(user.HashedPassword)
		if err != nil {
			return err
		}
		user.HashedPassword = hashedPassword
	}
	return s.repo.UpdateUser(user)
}

// 🟢 유저 삭제
func (s *UserService) DeleteUser(email string) error {
	return s.repo.DeleteUser(email)
}
