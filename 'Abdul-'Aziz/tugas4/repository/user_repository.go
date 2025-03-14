package repository

import (
	"errors"
	"tugas/tugas4/infrastructure/env"
	"tugas/tugas4/models"
)

type UserRepository interface {
    GetUser() (*models.User, error)
    SaveUser(user *models.User) error
}

type InMemoryUserRepo struct {
    user *models.User
}

func NewInMemoryUserRepo(env *env.EnvLoader) *InMemoryUserRepo {
    return &InMemoryUserRepo{
        user: &models.User{
            Name:            env.Get("NAMA"),
            Email:           env.Get("EMAIL"),
            Password:        env.Get("PASSWORD"),
            Vehicle:         "",
            Equipment:       []string{},
            Recommendations: []models.Recommendation{},
            Friends:         []models.Friend{},
        },
    }
}

func (r *InMemoryUserRepo) GetUser() (*models.User, error) {
    if r.user == nil {
        return nil, errors.New("user tidak ditemukan")
    }
    return r.user, nil
}

func (r *InMemoryUserRepo) SaveUser(user *models.User) error {
    r.user = user
    return nil
}