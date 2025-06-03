package models

import "time"

type User struct {
	UserID    int        `json:"user_id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Role      string     `json:"role"`
	Type      string     `json:"membership"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
