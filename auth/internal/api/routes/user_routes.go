package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/api/handler"
	"github.com/nibir30/go-microservices/auth/internal/constants"
	"github.com/nibir30/go-microservices/auth/internal/service"
)

func RegisterUserRoutes(router *gin.Engine, userService service.UserService, authService service.AuthService) {
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthenticationHandler(authService)

	userRoutes := router.Group(constants.UserRoutes)
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.POST("/", userHandler.CreateUser)
	}

	authRoutes := router.Group(constants.AuthRoutes)
	{
		authRoutes.POST("/login", authHandler.Login)
	}

}
