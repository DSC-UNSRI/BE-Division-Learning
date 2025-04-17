package main

import (
	"pertemuan05/config"
	"pertemuan05/database"
	"pertemuan05/routes"

	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.RoutesHandlers()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}