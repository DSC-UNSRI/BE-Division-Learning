package models

import "time"

type Menu struct {
	MenuID      int        `json:"menu_id"`
	Name        string     `json:"menu_name"`
	Description string     `json:"description"`
	Price       int        `json:"price"`
	ChefID      int        `json:"chef_id"`
	Category    string     `json:"category"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

var MenuQuery = `
	CREATE TABLE IF NOT EXISTS menus (
		menu_id INT AUTO_INCREMENT PRIMARY KEY,
		menu_name VARCHAR(30) NOT NULL,
		description VARCHAR(60),
		price INT NOT NULL,
		chef_id INT,
		category VARCHAR(30),
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (chef_id) REFERENCES chefs(chef_id)
	);`
