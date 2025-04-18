package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	db, err := gorm.Open(sqlite.Open("playlist.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
