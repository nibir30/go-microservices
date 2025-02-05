package bootstrap

import (
	"github.com/nibir30/go-microservices/auth/internal/repository"
	"github.com/nibir30/go-microservices/auth/internal/service"
	"gorm.io/gorm"
)

// Container holds all services and repositories
type Container struct {
	UserRepo    repository.UserRepository
	UserService service.UserService

	// Add other repositories and services as needed
}

// NewContainer initializes all repositories and services
func NewContainer(db *gorm.DB) *Container {
	return &Container{
		UserRepo:    repository.NewUserRepository(db),
		UserService: service.NewUserService(repository.NewUserRepository(db)),

		// Initialize more repositories and services here
	}
}
