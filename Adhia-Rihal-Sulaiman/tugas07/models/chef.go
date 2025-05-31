package models

import "time"

type Chef struct {
	ChefID     int        `json:"chef_id"`
	Name       string     `json:"chef_name"`
	Speciality string     `json:"speciality"`
	Experience int        `json:"experience"`
	Username   string     `json:"username"`
	Password   string     `json:"-"`
	Token      string     `json:"token"`
	Role       string     `json:"role"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var ChefQuery = `
	CREATE TABLE IF NOT EXISTS chefs (
		chef_id INT AUTO_INCREMENT PRIMARY KEY,
		chef_name VARCHAR(30) NOT NULL,
		speciality VARCHAR(30),
		experience INT,
		username VARCHAR(30) NOT NULL,
		password VARCHAR(100) NOT NULL,
		token VARCHAR(255) NOT NULL,
		role ENUM('old', 'new') NOT NULL DEFAULT 'new',
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`
