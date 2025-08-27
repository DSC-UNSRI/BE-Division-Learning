package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	Password       string `json:"passsword" gorm:"not null"`
	ProfilePicture string `json:"profile_picture" gorm:"default:'http://127.0.0.1:3000/assets/profile_pictures/auby.jpeg'"`
	Email          string `json:"email" gorm:"unique;not null"`
	Role           string `json:"role" gorm:"type:varchar(10);default:'user';not null"`
}

type UserLogin struct {
	gorm.Model
	Role           string `json:"role" gorm:"type:varchar(10);default:'user';not null"`
}

func (User) TableName() string {
    return "user"
}