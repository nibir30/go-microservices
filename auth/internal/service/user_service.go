package service

import (
	"time"

	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/model/common"
	"github.com/nibir30/go-microservices/auth/internal/repository"
	"github.com/nibir30/go-microservices/auth/internal/utils"
)

type UserService interface {
	GetAllUsers() ([]model.User, *common.CustomError)
	CreateUser(user *model.User) (*model.User, *common.CustomError)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]model.User, *common.CustomError) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, common.NewCustomError("Failed to get users", err.Error())

	}
	return users, nil

}

func (s *userService) CreateUser(user *model.User) (*model.User, *common.CustomError) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, common.NewCustomError("Invalid password", err.Error())
	}

	existingUser, err := s.userRepo.ExistsByUsername(user.Username)
	if err != nil {
		return nil, common.NewCustomError("Invalid username", err.Error())
	}

	if existingUser {
		return nil, common.ValidationError("Username already exists")
	}

	user.Password = hashedPassword

	// Set the registration date
	now := time.Now()
	user.RegistrationDate = &now

	// Call the repository to save the user
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, common.NewCustomError("Failed to create user", err.Error())
	}

	return user, nil
}
