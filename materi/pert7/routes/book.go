package routes

import (
    "be_pert5/controllers"
    "be_pert5/middleware"
    "be_pert5/utils"
    "net/http"
)

func BooksRoutes() {
    http.HandleFunc("/books", booksHandler)
    http.HandleFunc("/books/", booksHandlerWithID)
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

func booksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        withAuth(controllers.GetBooks)(w, r)
    case http.MethodPost:
        withAdminAuth(controllers.CreateBook)(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func booksHandlerWithID(w http.ResponseWriter, r *http.Request) {
    parts, err := utils.SplitPath(r.URL.Path)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    id := parts[2]
    
    switch r.Method {
    case http.MethodPatch:
        withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.UpdateBook(w, r, id)
        })(w, r)
    case http.MethodDelete:
        withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.DeleteBook(w, r, id)
        })(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}