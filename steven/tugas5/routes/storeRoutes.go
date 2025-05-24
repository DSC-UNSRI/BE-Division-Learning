package routes

import (
	"tugas5/controllers"
	"tugas5/utils"
	"net/http"
)

func StoreRoutes(){
	http.HandleFunc("/store", storeHandler)
	http.HandleFunc("/store/", storeHandlerWithID)
	http.HandleFunc("/store/auth", storeAuthHandler)
}

func storeHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		controllers.GetStore(w, r, id)
	case http.MethodPatch:
		controllers.UpdateStore(w, r, id)
	case http.MethodDelete:
		controllers.DeleteStore(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func storeAuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.AuthStore(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}