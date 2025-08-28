package database

import (
	"fmt"
	"os"
	"backend2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBLoad() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection established")
	DB = db
}

func DBMigrate() {
	if err := DB.Debug().AutoMigrate(&models.Event{}, models.User{}); err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migration completed")
}
