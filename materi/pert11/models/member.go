package models

type Member struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"type:varchar(100)"`
	Password   string    `json:"password" gorm:"type:varchar(100)"`
	Gender     string    `json:"gender" gorm:"type:enum('male', 'female')"`
	DivisionID int       `json:"division_id"`
	Division   *Division `json:"division" gorm:"foreignKey:DivisionID;references:ID"`
	Project    *Project  `json:"project" gorm:"foreignKey:MemberID;references:ID"`
}

func (*Member) TableName() string {
	return "member"
}
