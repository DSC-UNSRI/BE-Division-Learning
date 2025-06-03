package models

import "time"

type Challenge struct {
	ChallengeID int       `json:"challenge_id"`
	Question    string    `json:"question"`
	Answer      string    `json:"answer"`
	CreatedAt   time.Time `json:"created_at"`
}
