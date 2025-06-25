package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	UserType string `json:"user_type"`
}

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}