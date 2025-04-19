package models

type NilaiMataUang struct {
	ID             int     `json:"id"`
	NamaMataUang   string  `json:"nama_mata_uang"`
	Singkatan      string  `json:"singkatan"`
	ExchangeRate   float64 `json:"exchange_rate"`
}