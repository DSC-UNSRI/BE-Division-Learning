package routes

import (
    "be_pert5/controllers"
    "net/http"
)

func UserRoutes() {
    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/register", controllers.Register)
}