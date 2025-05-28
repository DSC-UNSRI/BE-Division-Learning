package models

import "time"

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Token     string     `json:"token"`
	Role	  string     `json:"role"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var UsersQuery = `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		token VARCHAR(255) NOT NULL,
		role ENUM('admin', 'user') NOT NULL DEFAULT 'user',
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`
