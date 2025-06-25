package models

import "time"

type Lecturer struct {
	LecturerID   string     `json:"lecturer_id"`
	LecturerName string     `json:"lecturer_name"`
	Password     string     `json:"-"`
	Token        string     `json:"token"`
	Role         string     `json:"role"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

var LecturerQuery = `
	CREATE TABLE IF NOT EXISTS lecturers (
		lecturer_id VARCHAR(3) PRIMARY KEY,
		lecturer_name VARCHAR(40) NOT NULL,
		password VARCHAR(100) NOT NULL,
		token VARCHAR(255) NOT NULL,
		role ENUM('old', 'new') NOT NULL DEFAULT 'new',
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`
