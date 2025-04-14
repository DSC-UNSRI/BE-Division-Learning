package usecase

import (
	"fmt"
	"tugas/tugas4/repository"
)

type TransportUseCase struct {
    repo repository.UserRepository
}

func NewTransportUseCase(repo repository.UserRepository) *TransportUseCase {
    return &TransportUseCase{repo: repo}
}

func (uc *TransportUseCase) UpdateVehicle(vehicle string) error {
    validVehicles := map[string]bool{
        "Kendaraan Pribadi": true,
        "Bus Kaleng":        true,
        "Nebeng":            true,
        "Travel":            true,
    }
    
    if !validVehicles[vehicle] {
        return fmt.Errorf("kendaraan tidak valid: %s", vehicle)
    }
    
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    user.Vehicle = vehicle
    return uc.repo.SaveUser(user)
}