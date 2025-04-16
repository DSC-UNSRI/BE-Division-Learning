package repositories

import (
	"playlist-app/database"
	"playlist-app/models"
)

func GetAllArtists() ([]models.Artist, error) {
	rows, err := database.DB.Query("SELECT id, name FROM artists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []models.Artist
	for rows.Next() {
		var artist models.Artist
		if err := rows.Scan(&artist.ID, &artist.Name); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}
