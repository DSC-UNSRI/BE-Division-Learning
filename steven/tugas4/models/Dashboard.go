package models

type Dashboard struct {
	Kendaraan        string
	Barang           []string
	Rekomendasi      map[string]string
	Teman            map[string]string
}