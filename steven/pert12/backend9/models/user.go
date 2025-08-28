package models

type User struct {
	ID         		int       	`json:"id" gorm:"primaryKey"`
	Name       		string    	`json:"name" gorm:"type:varchar(100)"`
	Email     		string   	`json:"email"`
	Password   		string    	`json:"-" gorm:"type:varchar(100)"`
	Role			string		`json:"role" gorm:"type:enum('admin','user')"`
	ProfilePicture	string		`json:"profile_picture" gorm:"type:varchar(100)"`
	Status			bool 		`json:"status" gorm:"type:boolean"`
}

func (*User) TableName() string {
	return "user"
}