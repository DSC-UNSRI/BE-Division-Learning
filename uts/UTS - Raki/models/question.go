package models

import "time"

type Question struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username,omitempty"`
	IsPromoted bool     `json:"is_promoted"`
}

type CreateQuestionRequest struct {
	Title   string `json:"title" validate:"required,min=5,max=255"`
	Content string `json:"content" validate:"required,min=10"`
}

type PromoteQuestionRequest struct {
	QuestionID int `json:"question_id" validate:"required"`
}