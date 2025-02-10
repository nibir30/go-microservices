package middleware

import (
	"github.com/gin-gonic/gin"
	jwtConfig "github.com/nibir30/go-microservices/auth/config/jwt"
	utils "github.com/nibir30/go-microservices/auth/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.ErrorResponse(c, "Unauthorized", "No token provided")
			c.Abort()
			return
		}

		jwtUser, err := jwtConfig.VerifyToken(tokenString)

		if err != nil {
			utils.ErrorResponse(c, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		c.Set("jwtUser", jwtUser)
		c.Next()
	}
}
