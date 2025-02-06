package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/utils"
)

// DefaultErrorHandler is the global error handler middleware
func DefaultErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Recover from panics and handle them gracefully
		defer func() {
			if err := recover(); err != nil {
				// Log the panic or any error
				fmt.Println("Recovered from panic:", err)

				// Send a generic error response
				utils.ErrorResponse(c, "Internal server error", err.(string), http.StatusInternalServerError)
			}
		}()
		// Continue processing the request
		c.Next()

	}
}
