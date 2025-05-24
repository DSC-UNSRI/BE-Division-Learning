package models

type Nasabah struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

