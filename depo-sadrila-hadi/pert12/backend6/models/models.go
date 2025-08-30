package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `json:"name"`
	Email          string         `gorm:"unique" json:"email"`
	Password       string         `json:"-"`
	Role           string         `json:"role"`
	ProfilePicture string         `json:"profile_picture"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

type Event struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Location  string         `json:"location"`
	Start     *time.Time     `json:"start"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}