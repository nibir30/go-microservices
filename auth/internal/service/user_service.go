package service

import (
	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.GetAllUsers()
}


func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}
