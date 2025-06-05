package routes

import (
	"net/http"
	"strings"

	"resepku/controllers"
	"resepku/middleware"
)

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Resep Makanan is running"))
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		controllers.Register(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		controllers.Login(w, r)
	})

	http.HandleFunc("/resep", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			middleware.Logger(controllers.GetAllResep)(w, r)
		case "POST":
			middleware.Logger(middleware.AuthMiddleware(middleware.AdminOnly(controllers.CreateResep)))(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/my-resep", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		middleware.AuthMiddleware(controllers.GetMyResep)(w, r)
	})

	http.HandleFunc("/resep/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/resep/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		r.URL.RawQuery = "id=" + id

		switch r.Method {
		case "GET":
			controllers.GetResepByID(w, r)
		case "PUT":
			controllers.UpdateResep(w, r)
		case "DELETE":
			controllers.DeleteResep(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/negara", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllNegara(w, r)
		case "POST":
			controllers.CreateNegara(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/negara/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/negara/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		r.URL.RawQuery = "id=" + id

		switch r.Method {
		case "GET":
			controllers.GetNegaraByID(w, r)
		case "PUT":
			controllers.UpdateNegara(w, r)
		case "DELETE":
			controllers.DeleteNegara(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/resep-negara", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		controllers.GetResepResultNegara(w, r)
	})

	http.HandleFunc("/resep-negara/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/resep-negara/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		r.URL.RawQuery = "id=" + id

		switch r.Method {
		case "GET":
			controllers.GetResepByNegaraID(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
