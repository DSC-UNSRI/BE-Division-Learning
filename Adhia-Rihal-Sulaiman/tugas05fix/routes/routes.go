package routes

import (
	"be_pert5/controllers"
	"be_pert5/utils"
	"net/http"
)

func ChefRoutes() {
	http.HandleFunc("/chefs", chefsHandler)
	http.HandleFunc("/chefs/", chefsHandlerWithID)
}

func MenuRoutes() {
	http.HandleFunc("/menus", menusHandler)
	http.HandleFunc("/menus/", menusHandlerWithID)
}

func chefsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetChefs(w, r)
	case http.MethodPost:
		controllers.CreateChef(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func chefsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodPatch:
		controllers.UpdateChef(w, r, id)
	case http.MethodDelete:
		controllers.DeleteChef(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func menusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetMenus(w, r)
	case http.MethodPost:
		controllers.CreateMenu(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func menusHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodPatch:
		controllers.UpdateMenu(w, r, id)
	case http.MethodDelete:
		controllers.DeleteMenu(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
