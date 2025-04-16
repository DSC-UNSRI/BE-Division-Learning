package models

import "fmt"

type Recommendation struct {
    Category string
    Content  string
}

var recommendations []Recommendation

func AddRecommendation() {
    var category, content string
    fmt.Print("Masukkan kategori rekomendasi: ")
    fmt.Scan(&category)
    fmt.Print("Masukkan isi rekomendasi: ")
    fmt.Scan(&content)

    recommendations = append(recommendations, Recommendation{Category: category, Content: content})
    fmt.Println("Rekomendasi berhasil ditambahkan.")
}
