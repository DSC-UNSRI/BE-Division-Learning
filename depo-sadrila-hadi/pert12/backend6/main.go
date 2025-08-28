package main

import (
	"log"
	"net/http"
	"nobar-backend/auth"
	"nobar-backend/database"
	"nobar-backend/handlers"
	"nobar-backend/middleware"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	auth.LoadJWTKey()
	database.Connect()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/register", handlers.Register).Methods("POST")
	api.HandleFunc("/login", handlers.Login).Methods("POST")
	api.HandleFunc("/logout", handlers.Logout).Methods("POST")

	api.HandleFunc("/event", handlers.GetEvents).Methods("GET")
	
	fs := http.FileServer(http.Dir("./assets/"))
    api.PathPrefix("/assets/").Handler(http.StripPrefix("/api/assets/", fs))


	s := api.PathPrefix("/").Subrouter()
	s.Use(middleware.JWTMiddleware)

	s.HandleFunc("/me", handlers.GetMe).Methods("GET")
	s.HandleFunc("/profile/{id}", handlers.UpdateProfile).Methods("PATCH")

	s.HandleFunc("/event", handlers.CreateEvent).Methods("POST")
	s.HandleFunc("/event/{id}", handlers.UpdateEvent).Methods("PATCH")
	s.HandleFunc("/event/{id}", handlers.DeleteEvent).Methods("DELETE")

	
	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
		gorillaHandlers.AllowCredentials(),
	)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}