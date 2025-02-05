package main

import (
	"github.com/nibir30/go-microservices/auth/bootstrap"
	"github.com/nibir30/go-microservices/auth/config"
	"gorm.io/gorm"
)



var (
	db *gorm.DB
)



func main() {
	app := bootstrap.InitializeApp()
	defer config.DisconnectDB(db)

	app.Router.Run(config.AuthServicePort)
}