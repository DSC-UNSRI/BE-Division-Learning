package routes

import (
	"be_pert5/controllers"
	"be_pert5/utils"
	"net/http"
)

func BooksRoutes(){
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", booksHandlerWithID)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetBooks(w, r)
	case http.MethodPost:
		controllers.CreateBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func booksHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]
	switch r.Method {
	case http.MethodPatch:
		controllers.UpdateBook(w, r, id)
	case http.MethodDelete:
		controllers.DeleteBook(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
