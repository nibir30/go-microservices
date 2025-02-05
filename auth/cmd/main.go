package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/auth/config"
	"gorm.io/gorm"
)


var (
	db *gorm.DB = config.ConnectDB()
)


func main() {
	defer config.DisconnectDB(db)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(config.AuthServicePort)
}