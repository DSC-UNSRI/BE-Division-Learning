package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Location string    `json:"location" gorm:"not null"`
	Start    time.Time `json:"start" gorm:"not null"`
	Cover    string    `json:"cover"`
}
