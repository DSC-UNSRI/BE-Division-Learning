package models

import "time"

type Address struct {
	AddressID  int    `json:"address_id"`
	LecturerID string `json:"lecturer_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var AddressQuery = `
	CREATE TABLE IF NOT EXISTS addresses (
		address_id INT AUTO_INCREMENT PRIMARY KEY,
		lecturer_id VARCHAR(3) NOT NULL,
		street VARCHAR(60),
		city VARCHAR(30),
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (lecturer_id) REFERENCES lecturers(lecturer_id)
	);`