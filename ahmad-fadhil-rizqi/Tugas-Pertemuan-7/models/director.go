package models

import "time"

type Director struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	PasswordHash string     `json:"-"`
	Token        string     `json:"token"`
	Role         string     `json:"role"`
	DeletedAt    *time.Time `json:"-"`
}

var DirectorsQuery = `
	CREATE TABLE IF NOT EXISTS directors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		token VARCHAR(255) UNIQUE,
		role ENUM('admin', 'user') NOT NULL DEFAULT 'user',
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`