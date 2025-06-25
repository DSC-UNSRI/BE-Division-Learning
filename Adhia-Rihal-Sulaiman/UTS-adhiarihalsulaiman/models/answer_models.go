package models

import "time"

type Answer struct {
	ID         int        `json:"answer_id"`
	QuestionID int        `json:"question_id"`
	UserID     int        `json:"user_id"`
	Content    string     `json:"content"`
	BestAnswer bool       `json:"best_answer"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var AnswerTableSchema = `
	CREATE TABLE IF NOT EXISTS answers (
		answer_id INT AUTO_INCREMENT PRIMARY KEY,
		question_id INT NOT NULL,
		user_id INT NOT NULL,
		content TEXT NOT NULL,
		best_answer BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (question_id) REFERENCES questions(question_id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
`
