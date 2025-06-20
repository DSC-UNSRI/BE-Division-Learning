package domain

import "time"

type UserTier string

const (
	TierFree    UserTier = "free"
	TierPremium UserTier = "premium"
)

type User struct {
	ID               int
	Name             string
	Email            string
	Password         string
	Tier             UserTier
	SecurityQuestion string
	SecurityAnswer   string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}
