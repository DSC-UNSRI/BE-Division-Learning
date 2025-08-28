package routes

import (
	"net/http"
	"uts/controllers"
	"uts/middleware"
)

func SetupRoutes() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/forgot-password", controllers.ForgotPassword)
	mux.HandleFunc("/reset-password", controllers.ResetPassword)

	mux.Handle("/users/me", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMyProfile)))
	mux.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetAllUsers)))
	mux.Handle("/users/find", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetUserByID)))

	mux.HandleFunc("/questions", controllers.GetAllQuestions)
	mux.HandleFunc("/questions/find", controllers.GetQuestionByID)

	mux.Handle("/questions/create", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateQuestion)))
	mux.Handle("/answers/create", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateAnswer)))


	return mux
}