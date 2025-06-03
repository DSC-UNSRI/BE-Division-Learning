package models

import "time"

type Answer struct {
	AnswerID   int       `json:"answer_id"`
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	Upvotes    int       `json:"upvotes"`
	Downvotes  int       `json:"downvotes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
