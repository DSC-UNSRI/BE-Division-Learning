package database

import (
	"log"
	"backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:@tcp(127.0.0.1:3306)/learn12?charset=utf8mb4&parseTime=True&loc=Local"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Event{})
	log.Println("Database migration completed.")
}