package models

import "time"

type Lecturer struct {
	LecturerID        int        `json:"lecturer_id"`
	LecturerName      string     `json:"lecturer_name"`
	Password  string     `json:"password"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var LecturerQuery = `
	CREATE TABLE IF NOT EXISTS lecturers (
		lecturer_id VARCHAR(3) PRIMARY KEY,
		lecturer_name VARCHAR(40) NOT NULL,
		password VARCHAR(30) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`