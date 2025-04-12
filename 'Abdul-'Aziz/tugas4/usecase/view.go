package usecase

import (
	"fmt"
	"strings"
	"tugas/tugas4/models"
	"tugas/tugas4/repository"
)

type ViewUseCase struct {
    repo repository.UserRepository
}

func NewViewUseCase(repo repository.UserRepository) *ViewUseCase {
    return &ViewUseCase{repo: repo}
}

func (uc *ViewUseCase) GetUserData() (*models.User, error) {
    return uc.repo.GetUser()
}

func (uc *ViewUseCase) GetFormattedData() (string, error) {
    user, err := uc.repo.GetUser()
    if err != nil {
        return "", err
    }

    var sb strings.Builder
    sb.WriteString("\n=== DATA PENDAFTARAN IFTAR ===\n")
    sb.WriteString(fmt.Sprintf("Nama: %s\n", user.Name))
    sb.WriteString(fmt.Sprintf("Email: %s\n", user.Email))
    sb.WriteString(fmt.Sprintf("Kendaraan: %s\n", user.Vehicle))

    sb.WriteString("\nPeralatan:\n")
    for i, item := range user.Equipment {
        sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, item))
    }

    sb.WriteString("\nRekomendasi:\n")
    for i, rec := range user.Recommendations {
        sb.WriteString(fmt.Sprintf("%d. [%s] %s\n", i+1, rec.Category, rec.Content))
    }

    sb.WriteString("\nDaftar Teman:\n")
    for i, friend := range user.Friends {
        sb.WriteString(fmt.Sprintf("%d. %s - %s\n", i+1, friend.Name, friend.Division))
    }

    return sb.String(), nil
}