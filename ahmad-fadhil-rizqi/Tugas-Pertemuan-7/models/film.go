package models

import "time"

type Film struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	DirectorID int        `json:"director_id"`
	DeletedAt  *time.Time `json:"-"`
}

var FilmsQuery = `
	CREATE TABLE IF NOT EXISTS films (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		director_id INT NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (director_id) REFERENCES directors(id) ON DELETE RESTRICT ON UPDATE CASCADE
	);
`