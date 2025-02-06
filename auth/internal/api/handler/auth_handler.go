package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/model/data"
	"github.com/nibir30/go-microservices/auth/internal/service"
	"github.com/nibir30/go-microservices/auth/internal/utils"
)

type AuthenticationHandler struct {
	authService service.AuthService
}

func NewAuthenticationHandler(authService service.AuthService) *AuthenticationHandler {
	return &AuthenticationHandler{authService: authService}
}

// @Summary Login
// @Description Login to the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body data.LoginRequest true "Login request"
// @Success 200 {object} data.LoginResponse "Login successful"
// @Router /auth/login [post]
func (h *AuthenticationHandler) Login(c *gin.Context) {
	var loginRequest data.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.ErrorResponse(c, "Invalid input", err.Error())
		return
	}

	loginResponse, err := h.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		utils.ErrorResponse(c, err.GetMessage(), err.ErrorDetails)
		return
	}

	utils.DataSuccessResponse(c, "Login successful", loginResponse)
}
