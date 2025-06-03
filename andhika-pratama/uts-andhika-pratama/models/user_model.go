package models

import "time"

type User struct {
	UserID    int        `json:"user_id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

var UserQuery = `
	CREATE TABLE IF NOT EXISTS users (
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role ENUM('user', 'admin') DEFAULT 'user',
		type ENUM('free', 'premium') DEFAULT 'free',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);
`