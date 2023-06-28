package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	// Get the database connection string from environment variable
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Connect to the database using the gorm package
	db, err = gorm.Open("mysql", dbConnectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Auto migrate the URL model
	db.AutoMigrate(&URL{})

}