package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/service"
	"github.com/nibir30/go-microservices/auth/internal/utils"
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
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userLog := user
	userLog.Password = "[MASKED]"
	log.Printf("Creating new user: %+v", userLog)


	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword
	now := time.Now()
	user.RegistrationDate = &now

	if err := h.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
