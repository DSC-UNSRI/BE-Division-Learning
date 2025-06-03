package models

import "time"

type Challenge struct {
	ChallengeID int       `json:"challenge_id"`
	UserID      int       `json:"user_id"`
	Question    string    `json:"question"`
	Answer      string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
}
