package database

import (
	"log"
	"nobar-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("nobar.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	DB = connection

	err = DB.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		log.Fatal("Failed to migrate database!")
	}
}