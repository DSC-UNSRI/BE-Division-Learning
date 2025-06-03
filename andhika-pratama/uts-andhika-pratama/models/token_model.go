package models

import "time"

type Token struct {
	TokenID    int       `json:"token_id"`
	TokenValue string    `json:"token_value"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}
