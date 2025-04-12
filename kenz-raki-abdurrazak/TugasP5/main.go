package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
)

func main() {
	config.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Seminar API is running")
	})

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
