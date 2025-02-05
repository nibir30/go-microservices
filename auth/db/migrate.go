package db

import (
	"log"

	"github.com/nibir30/go-microservices/auth/internal/model"
	"gorm.io/gorm"
)

// MigrateDB runs database migrations
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{}, // Add more models here
	)
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}
	log.Println("Database migration completed successfully")
}
