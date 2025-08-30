package models

type Event struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"id"`
	Location string `gorm:"column:location" json:"location"`
	Start    string `gorm:"column:start" json:"start"`
	Cover    string `gorm:"column:cover" json:"cover"`
}
