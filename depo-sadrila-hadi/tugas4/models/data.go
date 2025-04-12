package models

type Data struct {
	Kendaraan   []string
	Barang      []string
	Rekomendasi map[string][]string
	Teman       []Teman
}

type Teman struct {
	Nama   string
	Divisi string
}

func NewData() *Data {
	return &Data{
		Kendaraan:   []string{},
		Barang:      []string{},
		Rekomendasi: map[string][]string{},
		Teman:       []Teman{},
	}
}