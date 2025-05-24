package models

import "time"

type Programmer struct {
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Language          string     `json:"language"`
	YearsOfExperience int        `json:"years_of_experience"`
	UserID            int        `json:"user_id"`
	CreatedAt         *time.Time `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

var ProgrammerQuery = `
CREATE TABLE IF NOT EXISTS programmers (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	language VARCHAR(50) NOT NULL,
	years_of_experience INT NOT NULL,
	user_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);`
