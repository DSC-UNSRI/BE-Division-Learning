package routes

import (
	"uts/controllers"
	"uts/middlewares"

	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", middlewares.WithAuth(controllers.Logout))	
	http.HandleFunc("/password/forgot", controllers.InitiatePasswordReset)
	http.HandleFunc("/password/reset", controllers.PasswordReset)
}