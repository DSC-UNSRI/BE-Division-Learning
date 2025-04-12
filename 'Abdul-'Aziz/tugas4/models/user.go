package models

type User struct {
	Name            string
	Email           string
	Password        string
	Vehicle         string
	Equipment       []string
	Recommendations []Recommendation
	Friends         []Friend
}