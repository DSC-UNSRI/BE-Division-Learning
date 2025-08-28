package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int            `json:"id" gorm:"primarykey"`
	Name           string         `json:"name" gorm:"not null"`
	Password       string         `json:"-" gorm:"not null"`
	ProfilePicture string         `json:"profile_picture" gorm:"default:'http://127.0.0.1:3000/assets/profile_pictures/auby.jpeg'"`
	Email          string         `json:"email" gorm:"unique;not null"`
	Role           string         `json:"role" gorm:"type:varchar(10);default:'user';not null"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserLogin struct {
	ID        int            `json:"id" gorm:"primarykey"`
	Role      string         `json:"role" gorm:"type:varchar(10);default:'user';not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
