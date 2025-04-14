package usecase

import (
	"errors"

	"tugas/tugas4/repository"
)

type EquipmentUseCase struct {
    repo repository.UserRepository
}

func NewEquipmentUseCase(repo repository.UserRepository) *EquipmentUseCase {
    return &EquipmentUseCase{repo: repo}
}

func (uc *EquipmentUseCase) AddEquipment(items []string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    user.Equipment = append(user.Equipment, items...)
    return uc.repo.SaveUser(user)
}

func (uc *EquipmentUseCase) UpdateEquipment(index int, newItem string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    if index < 0 || index >= len(user.Equipment) {
        return errors.New("index peralatan tidak valid")
    }
    
    user.Equipment[index] = newItem
    return uc.repo.SaveUser(user)
}

