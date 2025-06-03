package models

import "time"

type Challenge struct {
	ChallengeID int       `json:"challenge_id"`
	UserID      int       `json:"user_id"`
	Question    string    `json:"question"`
	Answer      string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
}

var ChallengeQuery = `
	CREATE TABLE IF NOT EXISTS challenges (
		challenge_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		question VARCHAR(255) NOT NULL,
		answer VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`