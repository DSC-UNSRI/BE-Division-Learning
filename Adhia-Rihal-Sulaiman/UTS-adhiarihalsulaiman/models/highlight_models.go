package models

import "time"

type Highlight struct {
	ID          int        `json:"highlight_id"`
	UserID      int        `json:"user_id"`
	ContentType string     `json:"content_type"`
	ContentID   int        `json:"content_id"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}


var HighlightTableSchema = `
	CREATE TABLE IF NOT EXISTS highlights (
		highlight_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		content_type VARCHAR(50) NOT NULL,
		content_id INT NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP NULL DEFAULT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`