package models

import "time"

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	AuthKey   string     `json:"-"` 
	DeletedAt *time.Time `json:"-"`
}

var UsersQuery = `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE,
		auth_key VARCHAR(255) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`