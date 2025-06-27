package routes

import (
	"net/http"
	"UTS_BE/controllers"
	"UTS_BE/middleware"
)

func AuthRoutes() {
	http.HandleFunc("/auth/register", controllers.RegisterUser)
	http.HandleFunc("/auth/login", controllers.LoginUser)
	http.HandleFunc("/auth/reset-password", controllers.ResetPassword)
	http.HandleFunc("/auth/logout", middleware.WithAuth(controllers.LogoutUser))
	http.HandleFunc("/auth/profile", middleware.WithAuth(controllers.GetProfile))

}


