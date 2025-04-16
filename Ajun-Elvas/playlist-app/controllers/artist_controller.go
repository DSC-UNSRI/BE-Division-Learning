package controllers

import (
	"net/http"
	"playlist-app/repositories"

	"github.com/gin-gonic/gin"
)

func GetArtists(c *gin.Context) {
	artists, err := repositories.GetAllArtists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}
