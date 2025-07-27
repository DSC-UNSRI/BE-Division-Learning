package models

type Division struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Member     []*Member `json:"member" gorm:"foreignKey:DivisionID;references:ID"`
}

func (*Division) TableName() string {
	return "division"
}