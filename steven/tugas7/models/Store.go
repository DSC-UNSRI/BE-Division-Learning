package models

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Password string `json:"password"`
}

var StoreQuery = `
	CREATE TABLE stores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    owner VARCHAR(100),
    password VARCHAR(100),
	);`