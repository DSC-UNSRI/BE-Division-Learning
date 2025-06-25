package models

import "time"

type User struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Email            string    `json:"email" db:"email"`
	Password         string    `json:"-" db:"password"`
	Role             string    `json:"role" db:"role"`
	SecurityQuestion string    `json:"security_question" db:"security_question"`
	SecurityAnswer   string    `json:"-" db:"security_answer"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

var UserQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    security_question VARCHAR(255) NOT NULL,
    security_answer VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL
);
` 