package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string    `json:"name" gorm:"not null"`
	Email      string    `json:"email" gorm:"not null"`
	Password   string    `json:"-"`
	Role       string 	 `json:"role" gorm:"type:ENUM('user','admin');default:'user'"`
	ProfilePic string    `json:"profile_pic" gorm:"default:'/profile/default.png'"`
}
