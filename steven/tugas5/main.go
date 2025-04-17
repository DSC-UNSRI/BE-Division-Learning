package main

import (
	"tugas5/config"
	"tugas5/database"
	"tugas5/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.ProductsRoutes()
	routes.StoreRoutes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}