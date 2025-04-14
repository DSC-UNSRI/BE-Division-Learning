package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	UserID      int    `json:"user_id"`
}
