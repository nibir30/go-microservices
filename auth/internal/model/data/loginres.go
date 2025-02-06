package data

import "github.com/nibir30/go-microservices/auth/internal/model"

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
