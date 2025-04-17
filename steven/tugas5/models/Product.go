package models

import "time"

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    int `json:"price"`
	Stock    int     `json:"stock"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var ProductsQuery = `
	CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		price FLOAT,
		stock INT,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`