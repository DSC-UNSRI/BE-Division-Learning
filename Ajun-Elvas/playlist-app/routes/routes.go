package routes

import (
	"playlist-app/controllers"

	"github.com/gin-gonic/gin"

	"playlist-app/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	artistController *controllers.ArtistController,
	songController *controllers.SongController,
	userController *controllers.UserController,
) {
	api := router.Group("/api")

	// Public routes
	api.POST("/users", userController.Register)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Artist endpoints
		protected.POST("/artists", artistController.CreateArtist)
		protected.GET("/artists", artistController.GetAllArtists)
		protected.GET("/artists/:id", artistController.GetArtistByID)
		protected.PUT("/artists/:id", artistController.UpdateArtist)
		protected.DELETE("/artists/:id", artistController.DeleteArtist)
		protected.GET("/artists/:id/songs", songController.GetSongsByArtist)

		// Song endpoints
		protected.POST("/songs", songController.CreateSong)
		protected.GET("/songs", songController.GetAllSongs)
		protected.GET("/songs/:id", songController.GetSongByID)
		protected.PUT("/songs/:id", songController.UpdateSong)
		protected.DELETE("/songs/:id", songController.DeleteSong)

		// User info (by ID) & delete (dilindungi)
		protected.GET("/users/:id", userController.GetUserByID)
		protected.DELETE("/users/:id", userController.DeleteUser)
	}
}
