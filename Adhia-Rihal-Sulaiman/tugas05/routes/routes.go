package routes

import (
	"database/sql"
	"tugas05/controllers"
	"tugas05/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// Inisialisasi controller
	chefController := controllers.NewChefController(db)
	menuController := controllers.NewMenuController(db)

	// Group untuk chef
	chefRoutes := r.Group("/chefs")
	{
		// Endpoint registrasi chef
		chefRoutes.POST("/register", chefController.Create)

		// Endpoint login
		chefRoutes.POST("/login", chefController.Login)

		// Gunakan AuthMiddleware untuk rute yang memerlukan autentikasi
		chefRoutes.GET("", middleware.AuthMiddleware(), chefController.GetAll)
		chefRoutes.GET("/:id", middleware.AuthMiddleware(), chefController.GetByID)
		chefRoutes.PUT("/:id", middleware.AuthMiddleware(), chefController.Update)
		chefRoutes.DELETE("/:id", middleware.AuthMiddleware(), chefController.Delete)
	}

	// Group untuk menu
	menuRoutes := r.Group("/menus")
	{
		// Endpoint membuat menu baru
		menuRoutes.POST("", menuController.Create)

		// Endpoint mendapatkan semua menu
		menuRoutes.GET("", menuController.GetAll)

		// Endpoint pencarian menu
		menuRoutes.GET("/search", menuController.SearchMenus)

		// Endpoint mendapatkan menu berdasarkan ID
		menuRoutes.GET("/:id", menuController.GetByID)

		// Endpoint update menu
		menuRoutes.PUT("/:id", menuController.Update)

		// Endpoint hapus menu
		menuRoutes.DELETE("/:id", menuController.Delete)

		// Endpoint mendapatkan menu berdasarkan chef
		menuRoutes.GET("/chef/:chefId", menuController.GetMenusByChef)
	}

	// Contoh route kustom atau tambahan
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Restaurant Backend Service is running",
		})
	})
}
