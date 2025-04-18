package models

type MataUang struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	Singkatan  string  `json:"singkatan"`
	NilaiTukar float64 `json:"nilai_tukar"`
	NasabahID  int     `json:"nasabah_id"`
}
