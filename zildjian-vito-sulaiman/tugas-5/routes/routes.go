package routes

import (
	"database/sql"
	"net/http"
	"tugas-5/handlers"
	"tugas-5/repositories"
	"tugas-5/services"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	// Setup User routes
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup Programmer routes
	// programmerRepo := repositories.NewProgrammerRepository(db)
	// programmerService := services.NewProgrammerService(programmerRepo)
	// programmerHandler := handlers.NewProgrammerHandler(programmerService)

	// User endpoints
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Programmer endpoints
	// mux.HandleFunc("GET /programmers", programmerHandler.GetAllProgrammers)
	// mux.HandleFunc("GET /programmers/{id}", programmerHandler.GetProgrammerByID)
	// mux.HandleFunc("POST /programmers", programmerHandler.CreateProgrammer)
	// mux.HandleFunc("PUT /programmers/{id}", programmerHandler.UpdateProgrammer)
	// mux.HandleFunc("DELETE /programmers/{id}", programmerHandler.DeleteProgrammer)
}
