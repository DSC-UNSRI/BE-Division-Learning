package routes

import (
	"be_pert7/controllers"
	"be_pert7/middlewares"
	"be_pert7/utils"
	"net/http"
)

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, middleware.AuthMiddleware).ServeHTTP(w, r)
	}
}

func withOwnAuth(handler func(http.ResponseWriter, *http.Request, string), id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedControllerHandler := http.HandlerFunc(func(innerW http.ResponseWriter, innerR *http.Request) {
			handler(innerW, innerR, id)
		})

		middlewareChain := middleware.CourseOwnershipMiddleware(
			wrappedControllerHandler,
			id,
		)

		utils.ApplyMiddlewares(middlewareChain, middleware.AuthMiddleware).ServeHTTP(w, r)
	}
}

func withHeadAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.HeadMiddleware).ServeHTTP(w, r)
	}
}

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

}

func chefsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withHeadAuth(controllers.GetChefs)(w, r)
	case http.MethodPost:
		withHeadAuth(controllers.CreateChef)(w, r)
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
		withHeadAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetChefByID(w, r, id)
		})(w, r)
	case http.MethodPatch:
		withHeadAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateChef(w, r, id)
		})(w, r)
	case http.MethodDelete:
		withHeadAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteChef(w, r, id)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func menusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withAuth(controllers.GetMenus)(w, r)
	case http.MethodPost:
		withAuth(controllers.CreateMenu)(w, r)
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
		withOwnAuth(controllers.GetMenuByID, id)(w, r)
	case http.MethodPatch:
		withOwnAuth(controllers.UpdateMenu, id)(w, r)
	case http.MethodDelete:
		withOwnAuth(controllers.DeleteMenu, id)(w, r)
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
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetMenusByChef(w, r, chefID)
		})(w, r)
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
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetMenusByCategory(w, r, category)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
