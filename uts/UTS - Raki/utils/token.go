package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenerateOpaqueToken() string {
	return uuid.New().String()
}

func GetTokenExpiration() time.Time {
	return time.Now().Add(24 * time.Hour)
}