package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Location string `json:"location" gorm:"not null"`
	Start time.Time `json:"start" gorm:"not null"`
	Cover string `json:"cover" gorm:"default:'http://127.0.0.1:3000/assets/covers/hinamizawa.jpg'"`
}