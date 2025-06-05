package routes

import (
	"tugas5/controllers"
	"tugas5/utils"
	"tugas5/middleware"
	"net/http"
)

func ProductsRoutes(){
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productsHandlerWithID)
}

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        utils.ApplyMiddlewares(handler, middleware.AuthMiddleware).ServeHTTP(w, r)
    }
}

func withAdminAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.AdminMiddleware).ServeHTTP(w, r)
    }
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withAuth(controllers.GetProducts)(w, r)
	case http.MethodPost:
		withAdminAuth(controllers.CreateProduct)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func productsHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.GetProduct(w, r, id)
        })(w, r)
	case http.MethodPatch:
		withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.UpdateProduct(w, r, id)
        })(w, r)
	case http.MethodDelete:
		withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.DeleteProduct(w, r, id)
        })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


