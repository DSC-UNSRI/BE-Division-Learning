package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB = database
}
