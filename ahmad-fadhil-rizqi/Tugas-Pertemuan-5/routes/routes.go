package routes

import (
	"net/http"
	"strings"

	"Tugas-Pertemuan-5/controllers"
)

func Routes() {
	http.HandleFunc("/films", filmsHandler)
	http.HandleFunc("/films/", filmsHandlerWithID)

	http.HandleFunc("/directors", directorsHandler)
	http.HandleFunc("/directors/", directorsHandlerWithID)

	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/getfilmsbydirector/", filmsByDirectorHandler)
	http.HandleFunc("/getdirectorbyfilm/", directorByFilmHandler)

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

func directorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetDirectors(w, r)
	case http.MethodPost:
		controllers.CreateDirector(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func directorsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/directors/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		controllers.GetDirectorByID(w, r, id)
	case http.MethodPatch:
		controllers.UpdateDirector(w, r, id)
	case http.MethodDelete:
		controllers.DeleteDirector(w, r, id)
	default:
		http.Error(w, "Method not allowed for /directors/{id}", http.StatusMethodNotAllowed)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		controllers.Login(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func filmsByDirectorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	directorID := strings.TrimPrefix(r.URL.Path, "/getfilmsbydirector/")
	if directorID == "" || strings.Contains(directorID, "/") {
		http.NotFound(w, r)
		return
	}
	controllers.GetFilmsByDirectorID(w, r, directorID)
}

func directorByFilmHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	filmID := strings.TrimPrefix(r.URL.Path, "/getdirectorbyfilm/")
	if filmID == "" || strings.Contains(filmID, "/") {
		http.NotFound(w, r)
		return
	}
	controllers.GetDirectorByFilmID(w, r, filmID)
}