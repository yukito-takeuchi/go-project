package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() error {
	// Get database URL from environment variable
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Set GORM logger level based on environment
	logLevel := logger.Info
	if os.Getenv("ENV") == "production" {
		logLevel = logger.Error
	}

	// Connect to database
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully")

	// Auto migrate models here
	// Example: DB.AutoMigrate(&User{}, &Post{})

	return nil
}

// Example model structure (uncomment and modify as needed)
/*
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
}
*/
