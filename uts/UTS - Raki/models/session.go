package models

import "time"

type Session struct {
    Token     string    `json:"token"`
    UserID    int       `json:"user_id"`
    UserType  string    `json:"user_type"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
}