package usecase

import (
	"tugas/tugas4/repository"
)

type TransportUseCase struct {
    repo repository.UserRepository
}

func NewTransportUseCase(repo repository.UserRepository) *TransportUseCase {
    return &TransportUseCase{repo: repo}
}

func (uc *TransportUseCase) UpdateVehicle(vehicle string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    user.Vehicle = vehicle
    return uc.repo.SaveUser(user)
}