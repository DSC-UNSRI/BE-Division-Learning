package main

import (
	"be_pert7/config"
	"be_pert7/database"
	"be_pert7/routes"
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
	routes.AuthRoutes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}