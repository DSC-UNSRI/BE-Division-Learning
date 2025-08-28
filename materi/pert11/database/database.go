package database

import (
	"fmt"
	"os"
	"pert11/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FJakarta",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if dsn == "" {
		panic("DB_DSN not set in .env file")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
}

func Migrate() {
	if err := DB.Debug().AutoMigrate(&models.Event{}, &models.Division{}, &models.Member{}, &models.Project{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrate Successfuly")
}
