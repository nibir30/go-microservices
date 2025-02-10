package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const AuthServicePort = ":8001"

// ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load(".env")
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	} else {
		fmt.Println("Successfully connected to mysql database")
	}

	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	} else {
		fmt.Println("Successfully disconnected from mysql database")
	}
	dbSQL.Close()
}

// 	migrate -path "D:/Dev/Backend/go/go-microservices/auth/db/migrations" -database "mysql://root:13579@tcp(localhost:3306)/go_ms" up
