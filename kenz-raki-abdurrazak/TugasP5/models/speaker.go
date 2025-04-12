package models

type Speaker struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Expertise string `json:"expertise"`
	AuthKey   string `json:"auth_key"`
}
