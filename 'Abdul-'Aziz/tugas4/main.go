package main

import (
	"fmt"
	"tugas/tugas4/controller"
	"tugas/tugas4/infrastructure/env"
	"tugas/tugas4/infrastructure/input"
	"tugas/tugas4/repository"
	"tugas/tugas4/usecase"
)

func main() {
    envLoader := env.NewEnvLoader()
    if err := envLoader.Load(); err != nil {
        fmt.Println("Gagal memuat environment:", err)
        return
    }

    inputReader := input.NewInputReader()
    
    // Inisialisasi repository
    userRepo := repository.NewInMemoryUserRepo(envLoader)
    
    // Inisialisasi use case dengan interface
    authUC := usecase.NewAuthUseCase(userRepo)
    transportUC := usecase.NewTransportUseCase(userRepo)
    equipmentUC := usecase.NewEquipmentUseCase(userRepo)
    recommendationUC := usecase.NewRecommendationUseCase(userRepo)
    friendUC := usecase.NewFriendUseCase(userRepo)
    viewUC := usecase.NewViewUseCase(userRepo)

    menuController := controller.NewMenuController(
        authUC,
        transportUC,
        equipmentUC,
        recommendationUC,
        friendUC,
        viewUC,
        *inputReader,
    )

    // Autentikasi
    fmt.Print("Email: ")
    email := inputReader.ReadLine()
    fmt.Print("Password: ")
    password := inputReader.ReadLine()

    if ok, _ := authUC.Authenticate(email, password); !ok {
        fmt.Println("Autentikasi gagal")
        return
    }

    menuController.ShowMainMenu()
}