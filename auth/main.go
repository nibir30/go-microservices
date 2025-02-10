package main

import (
	"github.com/nibir30/go-microservices/auth/bootstrap"
	"github.com/nibir30/go-microservices/auth/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// @title Auth Service API
// @version 1.0
// @description This is the auth service
// @host localhost:8080
// @BasePath /auth

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := bootstrap.InitializeApp()

	defer config.DisconnectDB(db)

	app.Router.Run(config.AuthServicePort)
}
