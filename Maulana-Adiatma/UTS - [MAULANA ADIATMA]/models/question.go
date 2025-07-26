package models

type Question struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"` 
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}