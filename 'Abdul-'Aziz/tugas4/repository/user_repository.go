package repository

import (
	"tugas/tugas4/models"
)

type UserRepository interface {
    GetUser() (*models.User, error)
    SaveUser(user *models.User) error
}