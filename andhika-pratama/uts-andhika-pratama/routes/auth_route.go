package routes

import (
	"uts/controllers"

	"net/http"
)

func AuthRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/forget", controllers.InitiatePasswordReset)
}