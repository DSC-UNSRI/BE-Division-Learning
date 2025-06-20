package main

import (
	"database/sql"
	"net/http"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/internal/question"
	"uts-zildjianvitosulaiman/internal/user"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) http.Handler {

	// --- User Dependencies ---
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- Question Dependencies ---
	questionRepo := question.NewRepository(db)
	questionService := question.NewService(questionRepo)
	questionHandler := question.NewHandler(questionService)

	// ===============================================
	// ===              API ROUTES                 ===
	// ===============================================

	// --- User Routes ---
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	protectedUserRoutes := auth.AuthMiddleware(http.HandlerFunc(userHandler.GetMyProfile))
	mux.Handle("GET /users/me", protectedUserRoutes)

	// --- Question Routes ---
	mux.Handle("POST /questions", auth.AuthMiddleware(http.HandlerFunc(questionHandler.CreateQuestion)))
	mux.Handle("GET /questions", auth.AuthMiddleware(http.HandlerFunc(questionHandler.GetAllQuestions)))
	mux.Handle("GET /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.GetQuestionByID)))
	mux.Handle("PUT /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.UpdateQuestion)))
	mux.Handle("DELETE /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.DeleteQuestion)))
	return mux
}
