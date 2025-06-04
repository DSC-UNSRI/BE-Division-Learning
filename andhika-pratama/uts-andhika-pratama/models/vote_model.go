package models

import "time"

type Vote struct {
	VoteID     int        `json:"vote_id"`
	UserID     int        `json:"user_id"`
	TargetID   int    `json:"target_id"`
	TargetType string     `json:"target_type"`
	VoteType   string     `json:"vote_type"`
	CreatedAt  time.Time  `json:"-"`
}

var VoteQuery = `
	CREATE TABLE IF NOT EXISTS votes (
		vote_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		target_id INT NOT NULL,
		target_type ENUM('question', 'answer') NOT NULL,
		vote_type ENUM('up', 'down') NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (user_id, target_id, target_type),
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`
