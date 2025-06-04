package models

type resep struct {
	ID             int    `json:"id"`
	NamaResep      string `json:"nama_resep"`
	DeskripsiResep string `json:"deskripsi_resep"`
	BahanUtama     string `json:"bahan_utama"`
	WaktuMasak     string `json:"waktu_masak"`
	NegaraID       int    `json:"negara_id"`
}
