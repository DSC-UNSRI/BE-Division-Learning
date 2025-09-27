package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	ProfilePicture string `json:"profile_picture"`
}
