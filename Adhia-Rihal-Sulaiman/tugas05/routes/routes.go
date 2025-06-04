package routes

import (
	"net/http"
	"tugas05/controllers"
)

func SetupRoutes() {
	// Inisialisasi controller
	chefController := controllers.NewChefController()
	menuController := controllers.NewMenuController()

	// Rute untuk Chef
	http.HandleFunc("/chefs/register", chefController.Create)
	http.HandleFunc("/chefs/login", chefController.Login)
	http.HandleFunc("/chefs", chefController.GetAll)
	http.HandleFunc("/chefs/", chefController.GetByID)   // id parameter untuk GetByID lewat URL
	http.HandleFunc("/chefs/update", chefController.Update)
	http.HandleFunc("/chefs/delete", chefController.Delete)

	// Rute untuk Menu
	http.HandleFunc("/menu", menuController.GetAll)
	http.HandleFunc("/menus/create", menuController.Create)
	http.HandleFunc("/menus/update", menuController.Update)
	http.HandleFunc("/menus/delete", menuController.Delete)
	http.HandleFunc("/menus/search", menuController.SearchMenus)
	http.HandleFunc("/menus/chef", menuController.GetMenusByChef)

	// Tambahan endpoint /health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "message": "Restaurant Backend Service is running"}`))
	})
}
