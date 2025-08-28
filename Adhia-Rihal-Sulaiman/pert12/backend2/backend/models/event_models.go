package models

import (
	"time"
)

type Event struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	Location string    `json:"location" gorm:"not null"`
	Start    time.Time `json:"start" gorm:"not null"`
	Cover    string    `json:"cover"`
}
