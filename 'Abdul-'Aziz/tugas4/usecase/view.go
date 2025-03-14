package usecase

import (
	"tugas/tugas4/models"
	"tugas/tugas4/repository"
)

type ViewUseCase struct {
    repo repository.InMemoryUserRepo
}

func NewViewUseCase(repo repository.InMemoryUserRepo) *ViewUseCase {
	
    return &ViewUseCase{repo: repo}
}

func (uc *ViewUseCase) GetUserData() (*models.User, error) {
    return uc.repo.GetUser()
}