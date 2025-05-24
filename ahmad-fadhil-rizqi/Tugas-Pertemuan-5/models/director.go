package models

import "time"

type Director struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"-"`
}

var DirectorsQuery = `
	CREATE TABLE IF NOT EXISTS directors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`