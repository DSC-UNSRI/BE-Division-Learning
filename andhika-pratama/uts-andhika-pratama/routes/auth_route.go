package routes

import (
	"uts/controllers"
	"uts/middlewares"

	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)

	http.HandleFunc("/forget-password-initiate", controllers.InitiatePasswordReset)
	http.HandleFunc("/forget-password-reset", controllers.PasswordReset)

	http.HandleFunc("/logout", middlewares.WithAuth(controllers.Logout))
}