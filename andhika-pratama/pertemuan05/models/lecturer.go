package models

import "time"

type Lecturer struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	DeletedAt *time.Time `json:"deleted_at"`
}
