package middleware

import (
	"bytes"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/internal/constants"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var pathsToHideRequestBody = []string{constants.UserRoutes}

func shouldHideRequestBody(path string) bool {
	path = strings.TrimRight(path, "/")
	for _, p := range pathsToHideRequestBody {
		if path == p {
			return true
		}
	}
	return false
}

// Write implements the Writer interface and captures the response body
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogRequestAndResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.Request.Header.Get("User-Agent")
		// headers := c.Request.Header
		headers := ""
		timezone := c.Request.Header.Get("Timezone")

		start := time.Now()

		// Check if the request method is POST, PUT, or PATCH, which usually have a body
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {

			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("Error reading request body:", err)
			}

			// Restore the body to the request after reading
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			if c.Request.Method == "POST" && shouldHideRequestBody(c.Request.URL.Path) {
				log.Printf("Request: %s %s %s\nUser-Agent: %s\nTimezone: %s\nHeaders: %v\nRequest Body: %s",
					c.Request.Method, c.Request.URL.Path, c.ClientIP(), userAgent, timezone, headers, string(bodyBytes))

			} else {
				log.Printf("Request: %s %s %s\nUser-Agent: %s\nTimezone: %s\nHeaders: %v\nRequest Body: %s",
					c.Request.Method, c.Request.URL.Path, c.ClientIP(), userAgent, timezone, headers, string(bodyBytes))

			}

		} else {
			log.Printf("Request: %s %s %s\nUser-Agent: %s\nTimezone: %s\nHeaders: %v",
				c.Request.Method, c.Request.URL.Path, c.ClientIP(), userAgent, timezone, headers)
		}

		// Create a buffer to capture the response body
		var responseBody bytes.Buffer
		// Create a custom response writer that captures the response
		writer := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &responseBody,
		}
		// Replace the original c.Writer with the custom writer
		c.Writer = writer

		// Process the request
		c.Next()

		duration := time.Since(start)
		log.Printf("Response: %s %s %v", c.Request.Method, c.Request.URL.Path, duration)

		if responseBody.Len() > 0 {
			log.Printf("Response Body: %s", responseBody.String())
		}
	}
}
