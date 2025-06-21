package models

import "time"

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Password string `json:"password"`
	Token     string     `json:"token"`
	Role	  string     `json:"role"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var StoreQuery = `
	CREATE TABLE stores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    owner VARCHAR(100),
    password VARCHAR(100),
	token VARCHAR(255) NOT NULL,
	role ENUM('admin', 'user') NOT NULL DEFAULT 'user',
	deleted_at TIMESTAMP NULL DEFAULT NULL
	);`