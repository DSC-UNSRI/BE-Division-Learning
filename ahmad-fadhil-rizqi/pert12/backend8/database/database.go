package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend8/models"
)

var DB *gorm.DB

func DBLoad() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successful")
	DB = db
}

func DBMigrate() {
	if err := DB.Debug().AutoMigrate(&models.User{}, &models.Event{}); err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migration completed")
}
