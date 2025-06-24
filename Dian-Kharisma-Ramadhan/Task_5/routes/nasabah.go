package routes

import (
	"Task_5/controllers"
	"net/http"
)

func NasabahRoutes() {
	http.HandleFunc("/nasabah/login", controllers.LoginNasabah)
	http.HandleFunc("/nasabah/register", controllers.RegisterNasabah)
}
