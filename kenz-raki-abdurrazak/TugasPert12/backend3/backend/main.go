package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	cors "github.com/gorilla/handlers"
	"backend/database"
	"backend/handlers"
	"backend/middleware"
)

func main() {
	database.Init()


	r := mux.NewRouter()

	allowedOrigins := cors.AllowedOrigins([]string{"http://localhost:5173"})
	allowedMethods := cors.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	allowedHeaders := cors.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/api/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/logout", handlers.Logout).Methods("POST")
	r.HandleFunc("/api/event", handlers.GetEvents).Methods("GET")
	
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.JwtAuth)

	protectedRoutes.HandleFunc("/me", handlers.GetMe).Methods("GET")
	protectedRoutes.HandleFunc("/profile/{id}", handlers.UpdateProfile).Methods("PATCH")
	protectedRoutes.HandleFunc("/event", handlers.CreateEvent).Methods("POST")
	protectedRoutes.HandleFunc("/event/{id}", handlers.UpdateEvent).Methods("PATCH")
	protectedRoutes.HandleFunc("/event/{id}", handlers.DeleteEvent).Methods("DELETE")

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", cors.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r))
}