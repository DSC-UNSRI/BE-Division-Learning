package models

import "time"

type Question struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateQuestionPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type QuestionWithAnswers struct {
	Question
	Answers []Answer `json:"answers"`
}