package repositories

import (
	"house-scanner-backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository 인터페이스 정의
type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(email string) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository: UserRepository 인스턴스 생성
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUserByEmail: 이메일로 유저 검색
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser: 새로운 유저 생성
func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// UpdateUser: 유저 정보 업데이트
func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Model(&models.User{}).Where("email = ?", user.Email).Updates(user).Error
}

// DeleteUser: 유저 삭제
func (r *userRepository) DeleteUser(email string) error {
	return r.db.Where("email = ?", email).Delete(&models.User{}).Error
}
