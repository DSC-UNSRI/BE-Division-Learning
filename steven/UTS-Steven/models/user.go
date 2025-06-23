package models

import "time"

type User struct {
	ID                	int			`json:"id"`
	Name              	string		`json:"name"`
	Email             	string		`json:"email"`
	Password          	string		`json:"password"`
	Role              	string 		`json:"role"`
	Token				string     	`json:"token"`
	ResetToken        	string		`json:"reset_token"`
	ResetTokenExpire  	string		`json:"reset_token_expire"`
	DeletedAt 			*time.Time 	`json:"deleted_at"`
}

var UserQuery = `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		role ENUM('free', 'premium') NOT NULL DEFAULT 'free',
		reset_token VARCHAR(255) DEFAULT NULL,
		reset_token_expire DATETIME DEFAULT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`