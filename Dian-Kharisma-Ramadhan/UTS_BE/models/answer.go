package models

import "time"

type Answer struct {
	ID          int        `json:"id"`
	QuestionID  int        `json:"question_id"`
	UserID      int        `json:"user_id"`
	Content     string     `json:"content"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Highlight  	bool       `json:"highlight"`
}
var AnswerSchema = `
CREATE TABLE IF NOT EXISTS answers (
	id INT AUTO_INCREMENT PRIMARY KEY,
	question_id INT NOT NULL,
	user_id INT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL,
	FOREIGN KEY (question_id) REFERENCES questions(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);
`