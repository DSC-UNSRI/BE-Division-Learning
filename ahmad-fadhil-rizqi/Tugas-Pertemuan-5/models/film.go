package models

import "time"

type Film struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Director  string     `json:"director"`
	DeletedAt *time.Time `json:"-"`
}

var FilmsQuery = `
	CREATE TABLE IF NOT EXISTS films (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		director VARCHAR(100) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`