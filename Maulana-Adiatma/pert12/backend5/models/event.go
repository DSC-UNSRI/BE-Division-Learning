package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Cover		string 		`json:"cover" gorm: "not null"`
	Location 	string		`json:"location" gorm: "not null"`
	Start		time.Time	`json:"startpost"`
}