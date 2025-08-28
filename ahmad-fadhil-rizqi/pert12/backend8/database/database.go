package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"backend8/config"
	"backend8/models"
)

var DB *gorm.DB

func Init(cfg config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if os.Getenv("AUTO_MIGRATE") == "true" {
  	if err := DB.AutoMigrate(&models.User{}, &models.Event{}); err != nil {
    log.Fatalf("migrate failed: %v", err)
  	}}
}
