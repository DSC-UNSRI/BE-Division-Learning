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
	http.HandleFunc("/getmenusbychef/", menusByChefHandler)
	http.HandleFunc("/getmenusbycategory/", menusByCategoryHandler)
}

func AuthRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
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
	case http.MethodGet:
		controllers.GetChefByID(w, r, id)
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
	case http.MethodGet:
		controllers.GetMenuByID(w, r, id)
	case http.MethodPatch:
		controllers.UpdateMenu(w, r, id)
	case http.MethodDelete:
		controllers.DeleteMenu(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func menusByChefHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chefID := parts[2]
	switch r.Method {
	case http.MethodGet:
		controllers.GetMenusByChef(w, r, chefID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func menusByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category := parts[2]
	switch r.Method {
	case http.MethodGet:
		controllers.GetMenusByCategory(w, r, category)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}