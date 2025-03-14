package usecase

import (
	"errors"

	"tugas/tugas4/models"
	"tugas/tugas4/repository"
)

type FriendUseCase struct {
    repo repository.UserRepository
}

func NewFriendUseCase(repo repository.UserRepository) *FriendUseCase {
    return &FriendUseCase{repo: repo}
}

func (uc *FriendUseCase) AddFriend(name, division string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    user.Friends = append(user.Friends, models.Friend{
        Name:     name,
        Division: division,
    })
    return uc.repo.SaveUser(user)
}

func (uc *FriendUseCase) UpdateFriend(index int, name, division string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    if index < 0 || index >= len(user.Friends) {
        return errors.New("index teman tidak valid")
    }
    
    user.Friends[index] = models.Friend{
        Name:     name,
        Division: division,
    }
    return uc.repo.SaveUser(user)
}