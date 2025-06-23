package routes

import(
	"uts-gdg/controllers"
	"net/http"
)

func UserRoutes(){
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)
}


