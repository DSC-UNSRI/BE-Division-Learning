package controllers

import (
	"net/http"
	"playlist-app/models"
	"playlist-app/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	Repo *repositories.SongRepository
}

func (ctrl *SongController) CreateSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.Repo.Create(&song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, song)
}

func (ctrl *SongController) GetAllSongs(c *gin.Context) {
	songs, err := ctrl.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func (ctrl *SongController) GetSongByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	song, err := ctrl.Repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}
	c.JSON(http.StatusOK, song)
}

func (ctrl *SongController) GetSongsByArtist(c *gin.Context) {
	artistID, _ := strconv.Atoi(c.Param("id"))
	songs, err := ctrl.Repo.FindByArtistID(uint(artistID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func (ctrl *SongController) UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song.ID = uint(id)
	err := ctrl.Repo.Update(&song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, song)
}

func (ctrl *SongController) DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.Repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
