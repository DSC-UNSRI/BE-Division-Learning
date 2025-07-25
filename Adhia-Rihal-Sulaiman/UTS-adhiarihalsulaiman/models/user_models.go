package models

import "time"

type User struct {
	ID        int        `json:"user_id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var UserTableSchema = `
	CREATE TABLE IF NOT EXISTS users (
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role ENUM('user', 'admin') DEFAULT 'user',
		type ENUM('free', 'premium') DEFAULT 'free',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`
