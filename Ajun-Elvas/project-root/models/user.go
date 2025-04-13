package models

type Rekomendasi struct {
	Kategori string
	Isi      string
}

type Teman struct {
	Nama   string
	Divisi string
}

type User struct {
	Nama        string
	Email       string
	Kendaraan   string
	Barang      []string
	Rekomendasi []Rekomendasi
	Teman       []Teman
}
