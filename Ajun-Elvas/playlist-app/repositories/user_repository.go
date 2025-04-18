package repositories

import (
	"playlist-app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Register(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	err := r.DB.Where("name = ?", name).First(&user).Error
	return user, err
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
