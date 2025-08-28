package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("pert12S.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB = database
}
