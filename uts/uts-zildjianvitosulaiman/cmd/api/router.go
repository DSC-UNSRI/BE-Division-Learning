package main

import (
	"database/sql"
	"net/http"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/internal/user"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) http.Handler {

	// --- User Dependencies ---
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// ===============================================
	// ===              API ROUTES                 ===
	// ===============================================

	// --- User Routes ---
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	protectedUserRoutes := auth.AuthMiddleware(http.HandlerFunc(userHandler.GetMyProfile))
	mux.Handle("GET /users/me", protectedUserRoutes)
	return mux
}
