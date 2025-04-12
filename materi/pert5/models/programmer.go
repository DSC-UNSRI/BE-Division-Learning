package models

import "time"

type Programmer struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Language  string     `json:"language"`
	UserID    int        `json:"user_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var ProgrammersQuery = `
	CREATE TABLE IF NOT EXISTS programmers (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		language VARCHAR(100) NOT NULL,
		user_id INT NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
`
