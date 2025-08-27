package models

type Event struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Division []*Division `json:"division" gorm:"many2many:division_events;"`
}

func (*Event) TableName() string {
	return "event"
}