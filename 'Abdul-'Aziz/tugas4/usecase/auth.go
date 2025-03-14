package usecase

import (
	"tugas/tugas4/repository"
)

type AuthUseCase struct {
    repo repository.InMemoryUserRepo
}

func NewAuthUseCase(repo repository.InMemoryUserRepo) *AuthUseCase {
    return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Authenticate(email, password string) (bool, error) {
    user, err := uc.repo.GetUser()
    if err != nil {
        return false, err
    }
    return user.Email == email && user.Password == password, nil
}