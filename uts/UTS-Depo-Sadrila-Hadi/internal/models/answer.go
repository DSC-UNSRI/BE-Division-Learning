package models

import "time"

type Answer struct {
	ID         int64     `json:"id"`
	Body       string    `json:"body"`
	QuestionID int64     `json:"question_id"`
	UserID     int64     `json:"user_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
}

type AnswerPayload struct {
	Body string `json:"body"`
}