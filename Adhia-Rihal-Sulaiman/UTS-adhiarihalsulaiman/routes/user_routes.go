package routes

import (
	"net/http"
	"uts_adhia/controllers"
	"uts_adhia/middlewares"
	"uts_adhia/utils"
)

func UserRoutes() {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandlerWithID)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middlewares.WithAdminAuth(controllers.CreateUser)(w, r)
	case http.MethodGet:
		middlewares.WithAdminAuth(controllers.GetUsers)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		middlewares.WithAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetUserByID(w, r, id)
		})(w, r)
	case http.MethodPatch:
		middlewares.WithAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateUser(w, r, id)
		})(w, r)
	case http.MethodDelete:
		middlewares.WithAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteUser(w, r, id)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
