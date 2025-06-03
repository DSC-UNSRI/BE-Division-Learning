package models

import "time"

type Question struct {
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Upvotes    int       `json:"upvotes"`
	Downvotes  int       `json:"downvotes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
