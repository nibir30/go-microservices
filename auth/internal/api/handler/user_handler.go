package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @Summary Get all users
// @Schemes
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.User "List of all users"
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	log.Printf("initGetUsers")
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// @Summary Create a new user
// @Schemes
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} model.User "User created successfully"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	log.Printf("initCreateUser")

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
