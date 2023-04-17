package repositories

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(User User) error
	FindUserByID(in uint) (*User, error)
	FindUserByUsername(username string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// สร้าง User ใหม่
func (r *userRepository) CreateUser(User User) error {
	return r.db.Create(&User).Error
}

// ค้นหาด้วย ID
func (r *userRepository) FindUserByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ค้นหาด้วย Username
func (r *userRepository) FindUserByUsername(username string) (*User, error) {
	var user User
	result := r.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
