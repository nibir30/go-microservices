package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)
	// Initialize the migrate instance with the file-based migrations

	m, err := migrate.New(
		"file://db/migrations", // Path to your migration files
		dbURL,                   // Database connection URL
	)


	if err != nil {
		log.Fatal("Failed to initialize migrate instance:", err)
	}

	// Apply migrations (this will apply all new migrations)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	// Check if migration was successful
	log.Println("Database migration completed successfully")
}