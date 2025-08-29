package database

import "tugas12/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Event{})
}
