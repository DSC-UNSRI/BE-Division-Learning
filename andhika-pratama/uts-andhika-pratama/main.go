package main

import (
	"uts/config"

	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}