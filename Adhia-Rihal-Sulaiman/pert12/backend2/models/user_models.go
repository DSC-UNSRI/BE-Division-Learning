package models

type User struct {
	ID             uint   `json:"id" gorm:"primarykey"`
	Name           string `json:"name" gorm:"not null"`
	Password       string `json:"passsword" gorm:"not null"`
	ProfilePicture string `json:"profile_picture" gorm:"not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	Role           string `json:"role" gorm:"type:varchar(10);default:'user';not null"`
}

type UserLogin struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Role string `json:"role" gorm:"type:varchar(10);default:'user';not null"`
}

func (User) TableName() string {
	return "user"
}
