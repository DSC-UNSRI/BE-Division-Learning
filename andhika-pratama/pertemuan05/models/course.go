package models

import "time"

type Course struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	LecturerID int        `json:"lecturer_id"`
	Semester   int        `json:"semester"`
	Credit     int        `json:"credit"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
