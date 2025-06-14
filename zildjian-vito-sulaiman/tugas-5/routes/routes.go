package routes

import (
	"database/sql"
	"net/http"
	"tugas-5/handlers"
	"tugas-5/middleware"
	"tugas-5/repositories"
	"tugas-5/services"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	// Setup User routes
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup Programmer routes
	programmerRepo := repositories.NewProgrammerRepository(db)
	programmerService := services.NewProgrammerService(programmerRepo)
	programmerHandler := handlers.NewProgrammerHandler(programmerService)

	mux.HandleFunc("POST /login", userHandler.Login)

	// User endpoints
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", middleware.AuthMiddleware(userHandler.GetAllUsers))
	mux.HandleFunc("GET /users/{id}", middleware.AuthMiddleware(userHandler.GetUserByID))
	mux.HandleFunc("PUT /users/{id}", middleware.AuthMiddleware(userHandler.UpdateUser))
	mux.HandleFunc("DELETE /users/{id}", middleware.AuthMiddleware(userHandler.DeleteUser))

	// Programmer endpoints
	mux.HandleFunc("GET /programmers", middleware.AuthMiddleware(programmerHandler.GetAllProgrammers))
	mux.HandleFunc("GET /programmers/{id}", middleware.AuthMiddleware(programmerHandler.GetProgrammerByID))
	mux.HandleFunc("POST /programmers", middleware.AuthMiddleware(programmerHandler.CreateProgrammer))
	mux.HandleFunc("PUT /programmers/{id}", middleware.AuthMiddleware(programmerHandler.UpdateProgrammer))
	mux.HandleFunc("DELETE /programmers/{id}", middleware.AuthMiddleware(programmerHandler.DeleteProgrammer))
	mux.HandleFunc("GET /programmers/users/{id}", middleware.AuthMiddleware(programmerHandler.GetProgrammersByUserID))
	mux.HandleFunc("GET /programmers/users/{id}/count", middleware.AuthMiddleware(programmerHandler.CountProgrammersByUserID))

}
