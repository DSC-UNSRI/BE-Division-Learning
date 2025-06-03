package models

import "time"

type Token struct {
	TokenID    int       `json:"token_id"`
	UserID     int       `json:"user_id"`
	TokenValue string    `json:"token_value"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

var TokenQuery = `
	CREATE TABLE IF NOT EXISTS tokens (
		token_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		token_value VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`
