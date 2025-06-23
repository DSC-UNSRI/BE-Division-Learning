package models

import "time"

type Question struct {
	ID        int			`json:"id"`
	UserID    int			`json:"user_id"`
	Title     string		`json:"title"`
	Content   string		`json:"content"`
	CreatedAt *time.Time	`json:"created_at"`
	DeletedAt *time.Time 	`json:"deleted_at"`
}

var QuestionQuery = `
	CREATE TABLE IF NOT EXISTS questions (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`