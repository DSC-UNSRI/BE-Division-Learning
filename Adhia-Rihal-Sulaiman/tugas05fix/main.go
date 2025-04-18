package main

import (
	"be_pert5/config"
	"be_pert5/database"
	"be_pert5/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.ChefRoutes()
	routes.MenuRoutes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}