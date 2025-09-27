package models

type User struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	ProfilePicture string `json:"profile_picture" gorm:"type:varchar(255);default:'https://i.pravatar.cc/150'"`
	Name           string `json:"name" gorm:"type:varchar(100)"`
	Email          string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Password       string `json:"password" gorm:"type:varchar(100)"`
	Role           string `json:"role" gorm:"type:enum('admin','user');default:'user'"`
}

type UserLogin struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}

func (*User) TableName() string {
	return "user"
}
