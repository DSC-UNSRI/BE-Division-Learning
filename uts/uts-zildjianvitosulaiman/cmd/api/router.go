package main

import (
	"database/sql"
	"net/http"
	"uts-zildjianvitosulaiman/internal/user" // Sesuaikan nama modul
)

func RegisterRoutes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// --- User Dependencies ---
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// ===============================================
	// ===              API ROUTES                 ===
	// ===============================================

	// --- User Routes ---
	mux.HandleFunc("POST /register", userHandler.Register)

	return mux
}
