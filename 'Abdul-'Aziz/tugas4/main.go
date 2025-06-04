package main

import (
	"fmt"
	"tugas/tugas4/controller"
	"tugas/tugas4/infrastructure/env"
	"tugas/tugas4/infrastructure/input"
	"tugas/tugas4/models"
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
    
    // Inisialisasi repository (menggunakan penyimpanan file untuk data permanen)
    userRepo := repository.NewFileUserRepo(envLoader)
    
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
    
    // Tambahkan data sampel setelah autentikasi berhasil
    addSampleData(userRepo)

    menuController.ShowMainMenu()
}


func addSampleData(userRepo repository.UserRepository) {
	// Ambil data user yang sudah ada
	user, err := userRepo.GetUser()
	if err != nil {
		return
	}
	
	// Perbarui data user dengan data contoh jika diperlukan
	if user.Vehicle == "" {
		user.Vehicle = "Bus Kaleng"
	}
	
	// Tambahkan peralatan jika belum ada
	if len(user.Equipment) == 0 {
		user.Equipment = []string{
			"Sajadah",
			"Botol Air",
			"Masker",
		}
	}
	
	// Tambahkan rekomendasi jika belum ada
	if len(user.Recommendations) == 0 {
		user.Recommendations = []models.Recommendation{
			{Category: "Makanan", Content: "Kurma Medjool"},
			{Category: "Minuman", Content: "Air Kelapa"},
		}
	}
	
	// Tambahkan teman jika belum ada
	if len(user.Friends) == 0 {
		user.Friends = []models.Friend{
			{Name: "BAMBAN", Division: "FE"},
			{Name: "ABDUL", Division: "BE"},
			{Name: "RIZKY", Division: "UI/UX"},
		}
	}
	
	// Simpan perubahan
	userRepo.SaveUser(user)
}