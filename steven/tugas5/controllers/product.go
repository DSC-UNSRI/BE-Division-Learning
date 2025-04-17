package controllers

import (
	"tugas5/database"
	"tugas5/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, price, stock FROM products WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Success",
		"products": products,
	})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	err := r.ParseForm() 
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	product.Name = r.FormValue("name")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		http.Error(w, "Invalid stock format", http.StatusBadRequest)
		return
	}

	product.Price = price
	product.Stock = stock


	res, err := database.DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", product.Name, product.Price, product.Stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	product.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product Created",
		"product":    product,
	})
	
}

func GetProduct(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	row := database.DB.QueryRow("SELECT id, name, price, stock FROM products WHERE id = ? AND deleted_at IS NULL", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
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

	product := models.Product{}
	err := database.DB.QueryRow("SELECT id, name, price, stock FROM products WHERE id = ? AND deleted_at IS NULL", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
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