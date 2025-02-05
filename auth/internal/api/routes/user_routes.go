package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/api/handler"
	"github.com/nibir30/go-microservices/auth/internal/service"
)


func RegisterUserRoutes(router *gin.Engine, userService service.UserService) {
	userHandler := handler.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.POST("/", userHandler.CreateUser)
	}
}