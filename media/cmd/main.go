package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nibir30/go-microservices/media/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)


func main() {
	defer config.DisconnectDB(db)
	router := gin.Default()

	router.Run(config.MediaServicePort)
}