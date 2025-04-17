package main

import (
	"database/sql"
	"log"
	"restaurant-backend/config"
	"restaurant-backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDatabase()
	defer db.Close()

	r := gin.Default()

	routes.SetupRoutes(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
