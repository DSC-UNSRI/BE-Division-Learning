package controllers

import (
	"encoding/json"
	"net/http"
	"tugas5/database"
	"tugas5/models"
	"database/sql"
)

func AuthStore(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	owner := r.FormValue("owner")
	password := r.FormValue("password")

	if owner == "" || password == "" {
		http.Error(w, "owner or password cannot be empty", http.StatusBadRequest)
		return
	}

	var store models.Store
	err = database.DB.QueryRow("SELECT id, name, owner, password FROM stores WHERE owner = ? AND password = ? AND deleted_at IS NULL", owner, password).
		Scan(&store.ID, &store.Name, &store.Owner, &store.Password)
	if err != nil {
		http.Error(w, "owner or password is wrong", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login Success",
		"store":   store,
	})
}

func GetStores(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, owner, password FROM stores WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	stores := []models.Store{}
	for rows.Next() {
		store := models.Store{}
		rows.Scan(&store.ID, &store.Name, &store.Owner, &store.Password)
		stores = append(stores, store)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Success",
		"stores": stores,
	})
}

func CreateStore(w http.ResponseWriter, r *http.Request) {
	store := models.Store{}
	err := r.ParseForm() 
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	store.Name = r.FormValue("name")
	store.Owner = r.FormValue("owner")
	store.Password = r.FormValue("password")


	res, err := database.DB.Exec("INSERT INTO stores (name, owner, password) VALUES (?, ?, ?)", store.Name, store.Owner, store.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	store.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Store Created",
		"store":    store,
	})
	
}

func GetStore(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	row := database.DB.QueryRow("SELECT id, name, owner, password FROM stores WHERE id = ? AND deleted_at IS NULL", id)

	var store models.Store
	err := row.Scan(&store.ID, &store.Name, &store.Owner, &store.Password)
	if err != nil {
		http.Error(w, "Store not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success Found",
		"store": store,
	})
}

func UpdateStore(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	store := models.Store{}
	err := database.DB.QueryRow("SELECT id, name, owner, password FROM stores WHERE id = ? AND deleted_at IS NULL", id).Scan(&store.ID, &store.Name, &store.Owner, &store.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Store not found", http.StatusNotFound)
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

	name := r.FormValue("name")
	owner := r.FormValue("owner")
	password := r.FormValue("password")
	if name != "" {
		store.Name = name
	}
	if owner != "" {
		store.Owner = owner
	}
	if password != "" {
		store.Password = password
	}

	_, err = database.DB.Exec("UPDATE stores SET name = ?, owner = ?, password = ? WHERE id = ? AND deleted_at IS NULL", store.Name, store.Owner, store.Password, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Store Updated",
		"store":    store,
	})
}


func DeleteStore(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {		
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM stores WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "store not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE stores SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete store", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Store Deleted",
		"id":      id,
	})
}