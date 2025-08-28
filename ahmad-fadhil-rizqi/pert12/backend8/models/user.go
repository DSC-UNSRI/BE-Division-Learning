package models

type User struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"id"`
	Name           string `gorm:"column:name" json:"name"`
	Password       string `gorm:"column:password" json:"password"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture"`
	Email          string `gorm:"uniqueIndex;column:email" json:"email"`
	Role           string `gorm:"column:role" json:"role"`
	Status         *int   `gorm:"column:status" json:"status"`
}
