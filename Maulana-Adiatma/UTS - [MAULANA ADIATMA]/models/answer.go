package models

type Answer struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	UserID     int    `json:"user_id"`
	Username   string `json:"username"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
}
