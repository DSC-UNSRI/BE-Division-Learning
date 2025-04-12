package controller

import (
	"fmt"
	"strings"
	"tugas/tugas4/infrastructure/input"
	"tugas/tugas4/usecase"
)

type MenuController struct {
    authUseCase          *usecase.AuthUseCase
    transportUseCase     *usecase.TransportUseCase
    equipmentUseCase     *usecase.EquipmentUseCase
    recommendationUseCase *usecase.RecommendationUseCase
    friendUseCase        *usecase.FriendUseCase
    viewUseCase          *usecase.ViewUseCase
    inputReader          input.InputReader
}

func NewMenuController(
    auth *usecase.AuthUseCase,
    transport *usecase.TransportUseCase,
    equipment *usecase.EquipmentUseCase,
    recommendation *usecase.RecommendationUseCase,
    friend *usecase.FriendUseCase,
    view *usecase.ViewUseCase,
    reader input.InputReader,
) *MenuController {
    return &MenuController{
        authUseCase:          auth,
        transportUseCase:     transport,
        equipmentUseCase:     equipment,
        recommendationUseCase: recommendation,
        friendUseCase:        friend,
        viewUseCase:          view,
        inputReader:          reader,
    }
}

func (c *MenuController) ShowMainMenu() {
    for {
        fmt.Println("\n=== DASHBOARD IFTAR GDGOC ===")
        fmt.Println("1. Pilih Kendaraan")
        fmt.Println("2. Kelola Peralatan")
        fmt.Println("3. Kelola Rekomendasi")
        fmt.Println("4. Kelola Teman")
        fmt.Println("5. Lihat Semua Data")
        fmt.Println("6. Exit")
        fmt.Print("Pilih menu: ")

        switch c.inputReader.ReadInt(1, 6) {
        case 1:
            c.handleTransport()
        case 2:
            c.handleEquipment()
        case 3:
            c.handleRecommendation()
        case 4:
            c.handleFriend()
        case 5:
            c.handleViewData()
        case 6:
            fmt.Println("Terima kasih! Selamat Iftar!")
            return
        }
    }
}

func (c *MenuController) handleTransport() {
    fmt.Println("\nPilih Kendaraan:")
    options := []string{"Kendaraan Pribadi", "Bus Kaleng", "Nebeng", "Travel"}
    for i, opt := range options {
        fmt.Printf("%d. %s\n", i+1, opt)
    }
    fmt.Print("Pilihan: ")
    
    choice := c.inputReader.ReadInt(1, 4)
    vehicle := options[choice-1]
    
    if err := c.transportUseCase.UpdateVehicle(vehicle); err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Kendaraan berhasil diperbarui!")
    }
}

func (c *MenuController) handleEquipment() {
    for {
        fmt.Println("\nKelola Peralatan:")
        fmt.Println("1. Tambah Peralatan")
        fmt.Println("2. Update Peralatan")
        fmt.Println("3. Kembali")
        fmt.Print("Pilihan: ")

        switch c.inputReader.ReadInt(1, 3) {
        case 1:
            fmt.Print("Masukkan peralatan (pisahkan dengan koma): ")
            items := strings.Split(c.inputReader.ReadLine(), ",")
            var cleanItems []string
            for _, item := range items {
                if trimmed := strings.TrimSpace(item); trimmed != "" {
                    cleanItems = append(cleanItems, trimmed)
                }
            }
            if err := c.equipmentUseCase.AddEquipment(cleanItems); err != nil {
                fmt.Println("Error:", err)
            }
        case 2:
            user, _ := c.viewUseCase.GetUserData()
            if len(user.Equipment) == 0 {
                fmt.Println("Belum ada peralatan")
                break
            }
            fmt.Print("Pilih nomor peralatan: ")
            index := c.inputReader.ReadInt(1, len(user.Equipment)) - 1
            fmt.Print("Masukkan peralatan baru: ")
            newItem := c.inputReader.ReadLine()
            if err := c.equipmentUseCase.UpdateEquipment(index, newItem); err != nil {
                fmt.Println("Error:", err)
            }
        case 3:
            return
        }
    }
}

func (c *MenuController) handleRecommendation() {
    for {
        fmt.Println("\nKelola Rekomendasi:")
        fmt.Println("1. Tambah Rekomendasi")
        fmt.Println("2. Update Rekomendasi")
        fmt.Println("3. Kembali")
        fmt.Print("Pilihan: ")

        switch c.inputReader.ReadInt(1, 3) {
        case 1:
            fmt.Print("Masukkan kategori: ")
            kategori := c.inputReader.ReadLine()
            fmt.Print("Masukkan isi: ")
            isi := c.inputReader.ReadLine()
            if err := c.recommendationUseCase.AddRecommendation(kategori, isi); err != nil {
                fmt.Println("Error:", err)
            }
        case 2:
            user, _ := c.viewUseCase.GetUserData()
            if len(user.Recommendations) == 0 {
                fmt.Println("Belum ada rekomendasi")
                break
            }
            fmt.Print("Pilih nomor rekomendasi: ")
            index := c.inputReader.ReadInt(1, len(user.Recommendations)) - 1
            fmt.Print("Masukkan kategori baru: ")
            kategori := c.inputReader.ReadLine()
            fmt.Print("Masukkan isi baru: ")
            isi := c.inputReader.ReadLine()
            if err := c.recommendationUseCase.UpdateRecommendation(index, kategori, isi); err != nil {
                fmt.Println("Error:", err)
            }
        case 3:
            return
        }
    }
}

func (c *MenuController) handleFriend() {
    for {
        fmt.Println("\nKelola Teman:")
        fmt.Println("1. Tambah Teman")
        fmt.Println("2. Update Teman")
        fmt.Println("3. Kembali")
        fmt.Print("Pilihan: ")

        switch c.inputReader.ReadInt(1, 3) {
        case 1:
            fmt.Print("Masukkan nama teman: ")
            nama := c.inputReader.ReadLine()
            fmt.Print("Masukkan divisi: ")
            divisi := c.inputReader.ReadLine()
            if err := c.friendUseCase.AddFriend(nama, divisi); err != nil {
                fmt.Println("Error:", err)
            }
        case 2:
            user, _ := c.viewUseCase.GetUserData()
            if len(user.Friends) == 0 {
                fmt.Println("Belum ada teman")
                break
            }
            fmt.Print("Pilih nomor teman: ")
            index := c.inputReader.ReadInt(1, len(user.Friends)) - 1
            fmt.Print("Masukkan nama baru: ")
            nama := c.inputReader.ReadLine()
            fmt.Print("Masukkan divisi baru: ")
            divisi := c.inputReader.ReadLine()
            if err := c.friendUseCase.UpdateFriend(index, nama, divisi); err != nil {
                fmt.Println("Error:", err)
            }
        case 3:
            return
        }
    }
}

func (c *MenuController) handleViewData() {
    data, err := c.viewUseCase.GetFormattedData()
    if err != nil {
        fmt.Println("Gagal menampilkan data:", err)
        return
    }
    fmt.Println(data)
}