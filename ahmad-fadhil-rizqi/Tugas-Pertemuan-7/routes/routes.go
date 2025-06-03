package routes

import (
	"net/http"
	"path"
	"strings"

	"Tugas-Pertemuan-7/controllers"
	"Tugas-Pertemuan-7/middleware"
	"Tugas-Pertemuan-7/utils"
)

func Routes() {
	http.HandleFunc("/register/director", controllers.RegisterDirector)
	http.HandleFunc("/login/director", controllers.LoginDirector)

	http.HandleFunc("/films", filmsHandler)
	http.HandleFunc("/films/", filmsHandlerWithID)

	http.HandleFunc("/directors", directorsHandler)
	http.HandleFunc("/directors/", directorsHandlerWithID)

	http.HandleFunc("/getfilmsbydirector/", filmsByDirectorHandler)
	http.HandleFunc("/getdirectorbyfilm/", directorByFilmHandler)
}

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(http.HandlerFunc(handler), middleware.AuthMiddleware).ServeHTTP(w, r)
	}
}

func withAdminAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(http.HandlerFunc(handler), middleware.AuthMiddleware, middleware.AdminMiddleware).ServeHTTP(w, r)
	}
}

func withLogging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(http.HandlerFunc(handler), middleware.LoggingMiddleware).ServeHTTP(w, r)
	}
}

func filmsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withLogging(withAuth(controllers.GetFilms))(w, r)
	case http.MethodPost:
		withLogging(withAuth(controllers.CreateFilm))(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func filmsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	if id == "" || !strings.Contains(r.URL.Path, "/films/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetFilmByID(w, r, id)
		}))(w, r)
	case http.MethodPatch:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateFilm(w, r, id)
		}))(w, r)
	case http.MethodDelete:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteFilm(w, r, id)
		}))(w, r)
	default:
		http.Error(w, "Method not allowed for /films/{id}", http.StatusMethodNotAllowed)
	}
}

func directorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withLogging(withAuth(controllers.GetDirectors))(w, r)
	case http.MethodPost:
		withLogging(withAdminAuth(controllers.CreateDirector))(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func directorsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	if id == "" || !strings.Contains(r.URL.Path, "/directors/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetDirectorByID(w, r, id)
		}))(w, r)
	case http.MethodPatch:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateDirector(w, r, id)
		}))(w, r)
	case http.MethodDelete:
		withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteDirector(w, r, id)
		}))(w, r)
	default:
		http.Error(w, "Method not allowed for /directors/{id}", http.StatusMethodNotAllowed)
	}
}

func filmsByDirectorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	directorID := path.Base(r.URL.Path)
	if directorID == "" || !strings.Contains(r.URL.Path, "/getfilmsbydirector/") {
		http.NotFound(w, r)
		return
	}
	withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
		controllers.GetFilmsByDirectorID(w, r, directorID)
	}))(w, r)
}

func directorByFilmHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	filmID := path.Base(r.URL.Path)
	if filmID == "" || !strings.Contains(r.URL.Path, "/getdirectorbyfilm/") {
		http.NotFound(w, r)
		return
	}
	withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
		controllers.GetDirectorByFilmID(w, r, filmID)
	}))(w, r)
}