package database

import (
	"fmt"
	"os"
	"pert12/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)


  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil{
	panic("Database failed to connect")
  }
  fmt.Println("Database succes to connect")
  DB = db
}

func DBMigrate(){	
	err := DB.AutoMigrate(&models.User{}, &models.Event{})

	if err != nil{
		panic("Database Migration Failed")
	}
	fmt.Println("Database succes to migrate")
}