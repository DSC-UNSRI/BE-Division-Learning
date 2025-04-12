package controllers

import (
	"be_pert5/database"
	"be_pert5/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, author FROM books WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		book := models.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author)
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"books": books,
	})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := r.ParseForm() //pakai Multipart jika ada file 
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")

	//jika bukan form
	// err := json.NewDecoder(r.Body).Decode(&book)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	res, err := database.DB.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	book.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "book successfully created",
		"book":    book,
	})
	
}

func UpdateBook(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	book := models.Book{}
	err := database.DB.QueryRow("SELECT id, title, author FROM books WHERE id = ? AND deleted_at IS NULL", id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	if title != "" {
		book.Title = title
	}
	if author != "" {
		book.Author = author
	}

	_, err = database.DB.Exec("UPDATE books SET title = ?, author = ? WHERE id = ? AND deleted_at IS NULL", book.Title, book.Author, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book updated successfully",
		"book":    book,
	})
}

func DeleteBook(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {		
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE books SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Book deleted successfully",
		"id":      id,
	})
}

