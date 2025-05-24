package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/routes"
)
	

func main() {
	config.InitDB()

	r := routes.InitRoutes()

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}