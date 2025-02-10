package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/model"
	"github.com/nibir30/go-microservices/auth/internal/model/data"
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
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User "List of all users"
// @Router /users [get]
// @Security BearerAuth
// @in header
// @name Authorization
func (h *UserHandler) GetUsers(c *gin.Context) {
	log.Printf("initGetUsers")

	jwtUser, exists := c.Get("jwtUser")

	if !exists {
		utils.ErrorResponse(c, "Unauthorized", "No token provided")
		return
	}

	log.Printf(jwtUser.(data.JwtUser).Username)

	users, err := h.userService.GetAllUsers()

	if err != nil {
		utils.ErrorResponse(c, err.GetMessage(), err.ErrorDetails)
		return
	}
	utils.DataSuccessResponse(c, "Users fetched successfully", users)

}

// @Summary Create a new user
// @Schemes
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} model.User "User created successfully"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	log.Printf("initCreateUser")

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, "Invalid input", err.Error())
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		utils.ErrorResponse(c, err.GetMessage(), err.ErrorDetails)
		return
	}

	utils.DataSuccessResponse(c, "User created successfully", createdUser)
}
