package models

import "time"

type Forget struct {
	ID         int       `json:"forget_id"`
	UserID     int       `json:"user_id"`
	Question   string    `json:"question"`
	AnswerHash string    `json:"-"`
	CreatedAt  time.Time `json:"-"`
}

var ForgetTableSchema = `
	CREATE TABLE IF NOT EXISTS forgets (
		forget_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		question VARCHAR(255) NOT NULL,
		answer_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`
