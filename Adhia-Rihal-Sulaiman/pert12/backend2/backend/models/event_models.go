package models

import (
	"time"
)

type Event struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	Location string    `json:"location" gorm:"not null"`
	Start    time.Time `json:"start" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Cover    string    `json:"cover" gorm:"default:'http://127.0.0.1:3000/assets/cover/IMZ2.jpg'"`
}
