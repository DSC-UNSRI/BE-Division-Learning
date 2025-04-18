package models

import "time"

type Chef struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Speciality string `json:"speciality"`
	Experience int    `json:"experience"`
	Username   string `json:"username"`
	Password   string `json:"PASSWORD"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var ChefQuery = `
	CREATE TABLE IF NOT EXISTS chefs (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		speciality VARCHAR(30),
		experience VARCHAR(30),
		username VARCHAR(30) NOT NULL,
		password VARCHAR(30) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`