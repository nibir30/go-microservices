package service

import (
	jwtConfig "github.com/nibir30/go-microservices/auth/config/jwt"
	"github.com/nibir30/go-microservices/auth/internal/model/common"
	"github.com/nibir30/go-microservices/auth/internal/model/data"
	"github.com/nibir30/go-microservices/auth/internal/repository"

	passwordUtils "github.com/nibir30/go-microservices/auth/internal/utils/password"
)

type AuthService interface {
	Login(username, password string) (*data.LoginResponse, *common.CustomError)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(username, password string) (*data.LoginResponse, *common.CustomError) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {

		return nil, common.NewCustomError("User not found", err.Error())
	}

	valid, err := passwordUtils.VerifyPassword(password, user.Password)
	if err != nil {
		return nil, common.NewCustomError("Invalid password", err.Error())
	}

	if !valid {
		return nil, common.ValidationError("Invalid password")
	}

	token, err := jwtConfig.CreateToken(username)

	if err != nil {
		return nil, common.NewCustomError("Failed to create token", err.Error())
	}

	return &data.LoginResponse{User: *user, Token: token}, nil
}
