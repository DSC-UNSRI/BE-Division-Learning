package routes

import (
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/controllers"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/middleware"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/speakers", controllers.GetAllSpeakers).Methods("GET")
	r.HandleFunc("/speakers/{id}", controllers.GetSpeakerByID).Methods("GET")
	r.HandleFunc("/events", controllers.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", controllers.GetEventByID).Methods("GET")
	r.HandleFunc("/speakers/{id}/events", controllers.GetEventsBySpeakerID).Methods("GET")

	r.HandleFunc("/speakers", middleware.AuthMiddleware(controllers.CreateSpeaker)).Methods("POST")
	r.HandleFunc("/speakers/{id}", middleware.AuthMiddleware(controllers.UpdateSpeaker)).Methods("PUT")
	r.HandleFunc("/speakers/{id}", middleware.AuthMiddleware(controllers.DeleteSpeaker)).Methods("DELETE")
	r.HandleFunc("/events", middleware.AuthMiddleware(controllers.CreateEvent)).Methods("POST")
	r.HandleFunc("/events/{id}", middleware.AuthMiddleware(controllers.UpdateEvent)).Methods("PUT")
	r.HandleFunc("/events/{id}", middleware.AuthMiddleware(controllers.DeleteEvent)).Methods("DELETE")
	r.HandleFunc("/full-event", middleware.AuthMiddleware(controllers.CreateFullEvent)).Methods("POST")
	
	r.HandleFunc("/auth/login", controllers.LoginSpeaker).Methods("POST")

	return r
}