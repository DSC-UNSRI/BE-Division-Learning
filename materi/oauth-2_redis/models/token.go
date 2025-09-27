package models

type Token struct {
	RefreshToken string `json:"refresh_token"`
	ParentToken  string `json:"parent_token"`
	UserID       int    `json:"user_id"`
	Exp          int64  `json:"exp"`
}