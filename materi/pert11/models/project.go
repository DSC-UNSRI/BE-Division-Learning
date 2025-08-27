package models

type Project struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	MemberID int    `json:"member_id"`
	Member   *Member `json:"member" gorm:"foreignKey:MemberID;references:ID"`
}

func (*Project) TableName() string {
	return "project"
}