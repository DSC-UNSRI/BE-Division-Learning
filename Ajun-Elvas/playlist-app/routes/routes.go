package routes

import (
	"playlist-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/artists", controllers.GetArtists)
}
