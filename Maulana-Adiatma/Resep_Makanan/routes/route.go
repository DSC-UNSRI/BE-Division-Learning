package routes

import (
	"net/http"
	"percobaan3/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Resep Makanan is running"))
	})

	router.HandleFunc("/resep", controllers.GetAllResep).Methods("GET")
	router.HandleFunc("/resep", controllers.CreateResep).Methods("POST")
	router.HandleFunc("/resep/{id}", controllers.GetResepByID).Methods("GET")
	router.HandleFunc("/resep/{id}", controllers.UpdateResep).Methods("PUT")
	router.HandleFunc("/resep/{id}", controllers.DeleteResep).Methods("DELETE")

	router.HandleFunc("/negara", controllers.GetAllNegara).Methods("GET")
	router.HandleFunc("/negara", controllers.CreateNegara).Methods("POST")
	router.HandleFunc("/negara/{id}", controllers.GetNegaraByID).Methods("GET")
	router.HandleFunc("/negara/{id}", controllers.UpdateNegara).Methods("PUT")
	router.HandleFunc("/negara/{id}", controllers.DeleteNegara).Methods("DELETE")

	router.HandleFunc("/resep-negara", controllers.GetResepJoinNegara).Methods("GET")
	router.HandleFunc("/resep-negara/{id}", controllers.GetResepByNegaraID).Methods("GET")
	
	return router
}
