package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"tugas5/database"
	"tugas5/models"
	"tugas5/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, price, stock, store_id FROM products WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.StoreID)
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Success",
		"products": products,
	})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	if name == "" {
		http.Error(w, "Product name cannot be empty", http.StatusBadRequest)
		return
	}

	var exists bool
	
	err = database.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM products WHERE name = ? AND deleted_at IS NULL)`,name).Scan(&exists)

	if err != nil {
		http.Error(w, "Error checking product name", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Product name already exists", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil || price < 0 {
		http.Error(w, "Invalid price format or value", http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		http.Error(w, "Invalid stock format or value", http.StatusBadRequest)
		return
	}

	storeIDCtx := r.Context().Value(utils.StoreIDKey)
    if storeIDCtx == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
	
    storeIDstr, ok := storeIDCtx.(string)
    if !ok {
        http.Error(w, "Invalid store ID", http.StatusInternalServerError)
        return
    }

	storeID, err := strconv.Atoi(storeIDstr)
	if err != nil  {
		http.Error(w, "Invalid store id format or value", http.StatusBadRequest)
		return
	}

	// Simpan ke database
	res, err := database.DB.Exec("INSERT INTO products (name, price, stock, store_id) VALUES (?, ?, ?, ?)", name, price, stock, storeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()

	product := models.Product{
        ID:      int(id),
        Name:    name,
        Price:   price,
        Stock:   stock,
        StoreID: storeID,
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product Created",
		"product": product,
	})
}


func GetProduct(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	storeID := r.Context().Value(utils.StoreIDKey)
	if storeID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var product models.Product

	err := database.DB.QueryRow(`SELECT id, name, price, stock, store_id FROM products WHERE id = ? AND store_id = ? AND deleted_at IS NULL`,id, storeID).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.StoreID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success Found",
		"product": product,
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	storeIDCtx := r.Context().Value(utils.StoreIDKey)
    if storeIDCtx == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    storeID, ok := storeIDCtx.(string)
    if !ok {
        http.Error(w, "Invalid store ID", http.StatusInternalServerError)
        return
    }

	var product models.Product
	err := database.DB.QueryRow("SELECT id, name, price, stock, store_id FROM products WHERE id = ? AND deleted_at IS NULL", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.StoreID)

	if strconv.Itoa(product.StoreID) != storeID {
        http.Error(w, "Forbidden: You cannot update another store's product", http.StatusForbidden)
        return
    }

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
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
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	if priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil || price < 0 {
			http.Error(w, "Invalid or negative price", http.StatusBadRequest)
			return
		}
		product.Price = price
	}
	
	if stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil || stock < 0 {
			http.Error(w, "Invalid or negative stock", http.StatusBadRequest)
			return
		}
		product.Stock = stock
	}

	if name != "" {
		product.Name = name
	}

	_, err = database.DB.Exec("UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ? AND deleted_at IS NULL", product.Name, product.Price, product.Stock, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product Updated ",
		"product":  product,
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {		
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE products SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product Deleted",
		"id":      id,
	})
}