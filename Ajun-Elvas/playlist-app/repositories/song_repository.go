package repositories

import (
	"playlist-app/models"

	"gorm.io/gorm"
)

type SongRepository struct {
	DB *gorm.DB
}

func (r *SongRepository) Create(song *models.Song) error {
	return r.DB.Create(song).Error
}

func (r *SongRepository) FindAll() ([]models.Song, error) {
	var songs []models.Song
	err := r.DB.Preload("Artist").Find(&songs).Error
	return songs, err
}

func (r *SongRepository) FindByID(id uint) (models.Song, error) {
	var song models.Song
	err := r.DB.Preload("Artist").First(&song, id).Error
	return song, err
}

func (r *SongRepository) FindByArtistID(artistID uint) ([]models.Song, error) {
	var songs []models.Song
	err := r.DB.Where("artist_id = ?", artistID).Find(&songs).Error
	return songs, err
}

func (r *SongRepository) Update(song *models.Song) error {
	return r.DB.Save(song).Error
}

func (r *SongRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Song{}, id).Error
}
