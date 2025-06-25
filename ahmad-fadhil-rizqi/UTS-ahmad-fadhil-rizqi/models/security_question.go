package models

type SecurityQuestion struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	Question     string `json:"question"`
	HashedAnswer string `json:"-"`
}

var CreateSecurityQuestionsTableQuery = `
CREATE TABLE IF NOT EXISTS security_questions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    question TEXT NOT NULL,
    hashed_answer VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`