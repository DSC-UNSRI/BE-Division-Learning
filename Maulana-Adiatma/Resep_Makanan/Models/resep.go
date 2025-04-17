package models

type Resep struct {
	ID        		int    `json:"id"`
	NamaResep 		string `json:"nama_resep"`
	DeskripsiResep 	string `json:"description"`
}
