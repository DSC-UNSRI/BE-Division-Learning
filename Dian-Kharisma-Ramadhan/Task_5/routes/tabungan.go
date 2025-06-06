package routes

import (
    "Task_5/controllers"
    "Task_5/middleware"
    "Task_5/utils"
    "net/http"
)

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        middleware.AuthMiddleware(handler).ServeHTTP(w, r)
    }
}

func TabunganRoutes() {
    http.HandleFunc("/tabungan/all", withAuth(controllers.GetAllTabungan))
    http.HandleFunc("/tabungan/create", withAuth(controllers.CreateTabungan))
    http.HandleFunc("/tabungan/update/", withAuth(updateTabunganHandler))
    http.HandleFunc("/tabungan/delete/", withAuth(deleteTabunganHandler))
}

func updateTabunganHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := parts[2]

	controllers.UpdateTabungan(w, r, id)
}

func deleteTabunganHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    parts, err := utils.SplitPath(r.URL.Path)
    if err != nil || len(parts) < 3 {
        http.Error(w, "Invalid URL path", http.StatusBadRequest)
        return
    }
    id := parts[2]

    controllers.DeleteTabungan(w, r, id)
}
