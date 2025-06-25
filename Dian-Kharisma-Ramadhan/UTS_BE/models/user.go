package models

import "time"

type User struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"-"`
	Role       string     `json:"role"`
	Token      string     `json:"token"`
	SecretCode string     `json:"secret_code"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

var UserSchema = `
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	role ENUM('free','premium') DEFAULT 'free',
	token VARCHAR(255) UNIQUE,
	secret_code VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);
`
