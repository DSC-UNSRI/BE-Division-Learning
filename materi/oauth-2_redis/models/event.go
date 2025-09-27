package models

import "time"

type Event struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Location string    `json:"location" gorm:"type:varchar(100)"`
	Start    time.Time `json:"start"`
	Cover    string    `json:"cover" gorm:"type:varchar(255);default:'https://i.pravatar.cc/150'"`
}
