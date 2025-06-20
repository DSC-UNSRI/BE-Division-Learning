package main

import (
	"database/sql"
	"net/http"
	"uts-zildjianvitosulaiman/internal/answer"
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

	// --- Answer Dependencies ---
	answerRepo := answer.NewRepository(db)
	answerService := answer.NewService(answerRepo)
	answerHandler := answer.NewHandler(answerService)

	// ===============================================
	// ===              API ROUTES                 ===
	// ===============================================

	// --- User Routes ---
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	mux.HandleFunc("POST /forgot-password/request", userHandler.RequestPasswordReset)
	mux.HandleFunc("POST /forgot-password/verify", userHandler.VerifyAndResetPassword)

	protectedUserRoutes := auth.AuthMiddleware(http.HandlerFunc(userHandler.GetMyProfile))
	mux.Handle("GET /users/me", protectedUserRoutes)

	// --- Question Routes ---
	mux.Handle("POST /questions", auth.AuthMiddleware(http.HandlerFunc(questionHandler.CreateQuestion)))
	mux.Handle("GET /questions", auth.AuthMiddleware(http.HandlerFunc(questionHandler.GetAllQuestions)))
	mux.Handle("GET /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.GetQuestionByID)))
	mux.Handle("PUT /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.UpdateQuestion)))
	mux.Handle("DELETE /questions/{id}", auth.AuthMiddleware(http.HandlerFunc(questionHandler.DeleteQuestion)))

	// --- Answer Routes ---

	mux.Handle("POST /questions/{questionId}/answers", auth.AuthMiddleware(http.HandlerFunc(answerHandler.CreateAnswer)))
	mux.Handle("GET /questions/{questionId}/answers", auth.AuthMiddleware(http.HandlerFunc(answerHandler.GetAnswersForQuestion)))

	// Rute yang berhubungan dengan jawaban spesifik
	mux.Handle("PUT /answers/{id}", auth.AuthMiddleware(http.HandlerFunc(answerHandler.UpdateAnswer)))
	mux.Handle("DELETE /answers/{id}", auth.AuthMiddleware(http.HandlerFunc(answerHandler.DeleteAnswer)))

	return mux
}
