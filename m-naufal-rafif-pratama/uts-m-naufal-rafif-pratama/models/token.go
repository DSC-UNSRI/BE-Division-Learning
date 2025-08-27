package models

import "time"

type Token struct {
	Value     string    `db:"value"`
	UserID    int       `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
}

var TokenQuery = `
CREATE TABLE IF NOT EXISTS tokens (
    value VARCHAR(36) PRIMARY KEY,
    user_id INT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
` 