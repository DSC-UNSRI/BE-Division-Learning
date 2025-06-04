package routes

import (
	"net/http"
	"Task_5/controllers"
)

func RegisterNasabahRoutes() {
	http.HandleFunc("/nasabah/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateNasabah(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/nasabah/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAllNasabah(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/nasabah/detail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetNasabahByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/nasabah/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			controllers.UpdateNasabah(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/nasabah/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			controllers.DeleteNasabah(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/nasabah/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.LoginNasabah(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
