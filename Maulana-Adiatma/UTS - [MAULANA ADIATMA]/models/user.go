package models

type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"-"`
	Role             string `json:"role"`
	Token            string `json:"token"`
	CreatedAt        string `json:"created_at"`
	ResetToken       string `json:"-"`
	ResetTokenExpiry string `json:"-"`
}