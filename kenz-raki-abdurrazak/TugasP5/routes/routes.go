package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/controllers"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/speakers", controllers.GetAllSpeakers).Methods("GET")
	r.HandleFunc("/speakers/{id}", controllers.GetSpeakerByID).Methods("GET")
	r.HandleFunc("/speakers", controllers.CreateSpeaker).Methods("POST")
	r.HandleFunc("/speakers/{id}", controllers.UpdateSpeaker).Methods("PUT")
	r.HandleFunc("/speakers/{id}", controllers.DeleteSpeaker).Methods("DELETE")
	r.HandleFunc("/events", controllers.GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", controllers.GetEventByID).Methods("GET")
	r.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{id}", controllers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", controllers.DeleteEvent).Methods("DELETE")

	return r
}
