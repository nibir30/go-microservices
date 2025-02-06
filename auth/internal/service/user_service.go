package service

import (
	"time"

	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/repository"
	"github.com/nibir30/go-microservices/auth/internal/utils"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) (*model.User, error)
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


func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Set the registration date
	now := time.Now()
	user.RegistrationDate = &now

	// Call the repository to save the user
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
