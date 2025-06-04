package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas05/config"
	"tugas05/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.InitDB()

	// Setup routing
	routes.SetupRoutes()

	// Jalankan server HTTP
	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
