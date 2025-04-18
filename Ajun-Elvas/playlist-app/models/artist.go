package models

type Artist struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Username string `json:"username"`
	Password string `json:"password"`
}
