package repository

import (
	"github.com/nibir30/go-microservices/auth/internal/model"
	"gorm.io/gorm"
)


type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
