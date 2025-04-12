package models

import "time"

type Book struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var BooksQuery = `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(100) NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL
	);`
