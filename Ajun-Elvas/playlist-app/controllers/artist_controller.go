package controllers

import (
	"net/http"
	"playlist-app/models"
	"playlist-app/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	Repo *repositories.ArtistRepository
}

func (ctrl *ArtistController) CreateArtist(c *gin.Context) {
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.Repo.Create(&artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, artist)
}

func (ctrl *ArtistController) GetAllArtists(c *gin.Context) {
	artists, err := ctrl.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}

func (ctrl *ArtistController) GetArtistByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	artist, err := ctrl.Repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (ctrl *ArtistController) UpdateArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	artist.ID = uint(id)
	err := ctrl.Repo.Update(&artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (ctrl *ArtistController) DeleteArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.Repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
