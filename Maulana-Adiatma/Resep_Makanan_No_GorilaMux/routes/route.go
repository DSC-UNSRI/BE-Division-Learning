package routes

import (
	"net/http"
	"strings"
	"strconv"
	"resepku/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Resep Makanan is running"))
	})

	http.HandleFunc("/resep", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllResep(w, r)
		case "POST":
			controllers.CreateResep(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
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
		id := r.URL.Query().Get("id")
		if id == "" {
			controllers.GetResepResultNegara(w, r)
		} else {
			if _, err := strconv.Atoi(id); err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			controllers.GetResepByNegaraID(w, r)
		}
	})
}