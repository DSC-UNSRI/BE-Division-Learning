package routes

import (
	"net/http"
	"UTS_BE/controllers"
)

func AuthRoutes() {
	http.HandleFunc("/auth/register", controllers.RegisterUser)
	http.HandleFunc("/auth/login", controllers.LoginUser)
	http.HandleFunc("/auth/reset-password", controllers.ResetPassword)
}


