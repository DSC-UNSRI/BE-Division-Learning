package models

import "time"

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	PasswordHash string `json:"-"`
	Token string `json:"-"`
	TokenExpiresAt *time.Time `json:"-"`
	Tier string `json:"tier"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

var CreateUsersTableQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    token VARCHAR(255) UNIQUE NULL,
	token_expires_at TIMESTAMP NULL,
    tier ENUM('free', 'premium') NOT NULL DEFAULT 'free',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);`
