package routes

import (
	"uts/controllers"

	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	// http.HandleFunc("/forget", controllers.ForgetPassword)
}