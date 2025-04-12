package models

import "time"

type Lecturer struct {
	ID        int        `json:"lecturer_id"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var LecturerQuery = `
	CREATE TABLE IF NOT EXISTS lecturers (
		lecturer_id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(40) NOT NULL,
		password VARCHAR(30) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`