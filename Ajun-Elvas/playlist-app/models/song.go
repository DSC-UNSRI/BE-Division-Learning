package models

type Song struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ArtistID int    `json:"artist_id"`
}
