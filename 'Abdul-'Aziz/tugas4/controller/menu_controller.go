package controller

import (
	"fmt"
	"strings"

	"tugas/tugas4/infrastructure/input"
	"tugas/tugas4/usecase"
)

type MenuController struct {
    authUseCase         *usecase.AuthUseCase
    transportUseCase    *usecase.TransportUseCase
    equipmentUseCase    *usecase.EquipmentUseCase
    recommendationUseCase *usecase.RecommendationUseCase
    friendUseCase       *usecase.FriendUseCase
    viewUseCase         *usecase.ViewUseCase
    inputReader         input.InputReader
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
   
        
        switch choice {
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

// Implementasi serupa untuk handleRecommendation, handleFriend, dan handleViewData