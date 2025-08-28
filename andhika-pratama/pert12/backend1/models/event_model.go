package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID        int            `json:"id" gorm:"primarykey"`
	Location  string         `json:"location" gorm:"not null"`
	Start     time.Time      `json:"start" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Cover     string         `json:"cover" gorm:"default:'http://127.0.0.1:3000/assets/covers/hinamizawa.jpg'"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
