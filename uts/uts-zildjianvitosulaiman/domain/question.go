package domain

import "time"

type Question struct {
	ID        int
	Title     string
	Body      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
