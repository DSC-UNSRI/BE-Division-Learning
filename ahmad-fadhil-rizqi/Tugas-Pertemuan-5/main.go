package main

import (
	"Tugas-Pertemuan-5/config"
	"Tugas-Pertemuan-5/database"
	"Tugas-Pertemuan-5/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.Routes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}