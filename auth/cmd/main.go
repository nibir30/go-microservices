package main

import (
	"github.com/nibir30/go-microservices/auth/config"
	"github.com/nibir30/go-microservices/auth/internal/bootstrap"
	"gorm.io/gorm"
)



var (
	db *gorm.DB = config.ConnectDB()
)


func main() {
	app := bootstrap.InitializeApp()
	defer config.DisconnectDB(db)
	
	app.Router.Run(config.AuthServicePort)
}