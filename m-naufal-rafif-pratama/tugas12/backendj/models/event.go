package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Location    string `json:"location"`
	Start       string `json:"start"`
	Cover       string `json:"cover"`
}
