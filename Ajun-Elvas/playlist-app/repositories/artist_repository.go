package repositories

import (
	"playlist-app/models"

	"gorm.io/gorm"
)

type ArtistRepository struct {
	DB *gorm.DB
}

func (r *ArtistRepository) Create(artist *models.Artist) error {
	return r.DB.Create(artist).Error
}

func (r *ArtistRepository) FindAll() ([]models.Artist, error) {
	var artists []models.Artist
	err := r.DB.Find(&artists).Error
	return artists, err
}

func (r *ArtistRepository) FindByID(id uint) (models.Artist, error) {
	var artist models.Artist
	err := r.DB.First(&artist, id).Error
	return artist, err
}

func (r *ArtistRepository) Update(artist *models.Artist) error {
	return r.DB.Save(artist).Error
}

func (r *ArtistRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Artist{}, id).Error
}
