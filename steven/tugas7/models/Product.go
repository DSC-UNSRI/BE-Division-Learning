package models

import "time"

type Product struct {
	ID       int     	`json:"id"`
	Name     string  	`json:"name"`
	Price    int 		`json:"price"`
	Stock    int     	`json:"stock"`
	StoreID	 int 		`json:"store_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var ProductsQuery = `
	CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		price FLOAT,
		stock INT,
		store_id INT,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
	);`