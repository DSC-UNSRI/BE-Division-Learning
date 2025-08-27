package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	ProfilePicture string `json:"profile_picture"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Event struct {
	gorm.Model
	Location string `json:"location"`
	Start    string `json:"start"`
	Cover    string `json:"cover"`
}