package routes

import (
	"net/http"
	"strings"

	"Tugas-Pertemuan-5/controllers"
)

func Routes() {
	http.HandleFunc("/films", filmsHandler)
	http.HandleFunc("/films/", filmsHandlerWithID)

	http.HandleFunc("/login", loginHandler)
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetFilms(w, r)
	case http.MethodPost:
		controllers.CreateFilm(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func filmsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/films/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		controllers.GetFilmByID(w, r, id)
	case http.MethodPatch:
		controllers.UpdateFilm(w, r, id)
	case http.MethodDelete:
		controllers.DeleteFilm(w, r, id)
	default:
		http.Error(w, "Method not allowed for /films/{id}", http.StatusMethodNotAllowed)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		controllers.Login(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}