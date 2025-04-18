package main

import (
	"log"
	"playlist-app/controllers"
	"playlist-app/database"
	"playlist-app/repositories"
	"playlist-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi DB dari package database
	if err := database.Init(); err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Repositories
	artistRepo := &repositories.ArtistRepository{DB: database.DB}
	songRepo := &repositories.SongRepository{DB: database.DB}
	userRepo := &repositories.UserRepository{DB: database.DB}

	// Controllers
	artistCtrl := &controllers.ArtistController{Repo: artistRepo}
	songCtrl := &controllers.SongController{Repo: songRepo}
	userCtrl := &controllers.UserController{Repo: userRepo}

	// Setup router
	r := gin.Default()
	routes.SetupRoutes(r, artistCtrl, songCtrl, userCtrl)

	// Run server
	r.Run(":8080")
}
