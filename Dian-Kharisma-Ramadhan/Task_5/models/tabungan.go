package models

type Tabungan struct {
	ID            int     `json:"id"`
	NamaMataUang  string  `json:"nama_mata_uang"`
	Singkatan     string  `json:"singkatan"`
	Saldo         int64   `json:"saldo"`
}