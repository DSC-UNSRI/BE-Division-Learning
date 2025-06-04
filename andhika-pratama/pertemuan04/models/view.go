package models

import "fmt"

func ViewAll() {
    fmt.Println("\nData User Iftar:")
    fmt.Println("Kendaraan:", selectedVehicle)
    
    fmt.Println("\nBarang Bawaan:")
    for _, item := range items {
        fmt.Println("-", item)
    }

    fmt.Println("\nRekomendasi:")
    for _, rec := range recommendations {
        fmt.Printf("%s: %s\n", rec.Category, rec.Content)
    }

    fmt.Println("\nTeman yang ikut:")
    for _, friend := range friends {
        fmt.Printf("%s - %s\n", friend.Name, friend.Divisi)
    }
}
