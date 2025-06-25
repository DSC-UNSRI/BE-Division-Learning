package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas7/config"
	"tugas7/database"
	"tugas7/routes"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.SetupRoutes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}