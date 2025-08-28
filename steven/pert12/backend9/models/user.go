package models

type User struct {
	ID         		int       	`json:"id" gorm:"primaryKey"`
	Name       		string    	`json:"name" gorm:"type:varchar(100)"`
	Email     		string   	`json:"email"`
	Password   		string    	`json:"-" gorm:"type:varchar(100)"`
	Role			string		`json:"role" gorm:"type:enum('admin','user')default:'user'"`
	ProfilePicture	string		`json:"profile_picture" gorm:"type:varchar(100)default:'https://i.pravatar.cc/150'"`
	Status			*bool 		`json:"status" gorm:"type:tinyint(1)default:null"`
}

func (*User) TableName() string {
	return "user"
}