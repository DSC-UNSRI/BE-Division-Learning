package models

import "time"

type Answer struct {
	ID          int       `json:"id"`
	QuestionID  int       `json:"question_id"`
	UserID      int       `json:"user_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Username    string    `json:"username,omitempty"`
}

type CreateAnswerRequest struct {
	Content string `json:"content" validate:"required,min=10"`
}