package models

import "time"

type Event struct {
	ID        int       `json:"id"    gorm:"primaryKey"`
	Location  string    `json:"name"  gorm:"type:varchar(100)"`
	Start     time.Time `json:"email" gorm:"type:dateTime"`
	Cover     string    `json:"-"     gorm:"type:varchar(100)"`
}

func (*Event) TableName() string {
	return "event"
}