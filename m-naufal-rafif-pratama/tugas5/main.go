package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas5/config"
	"tugas5/database"
	"tugas5/routes"
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