package domain

import "time"

type Answer struct {
	ID         int
	Body       string
	UserID     int
	QuestionID int
	Upvotes    int
	Downvotes  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
