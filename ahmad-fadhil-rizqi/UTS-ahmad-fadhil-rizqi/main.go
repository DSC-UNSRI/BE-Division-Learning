package main

import (
	"UTS-Ahmad-Fadhil-Rizqi/config"
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.Routes()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server berjalan di http://localhost:%s/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}