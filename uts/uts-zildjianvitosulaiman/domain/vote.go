package domain

import "time"

type VoteType int

const (
	VoteTypeUp   VoteType = 1
	VoteTypeDown VoteType = -1
)

type Vote struct {
	ID        int
	UserID    int
	AnswerID  int
	Type      VoteType
	CreatedAt time.Time
}
