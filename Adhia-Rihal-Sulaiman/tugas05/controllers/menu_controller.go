package controllers  

import (  
	"database/sql"  
	"fmt"  
	"net/http"  
	"strconv"  

	"restaurant-backend/models"  

	"github.com/gin-gonic/gin"  
)  

type MenuController struct {  
	db *sql.DB  
}  

func NewMenuController(db *sql.DB) *MenuController {  
	return &MenuController{db: db}  
}  

func (c *MenuController) Create(ctx *gin.Context) {  
	var menu models.Menu  
	if err := ctx.ShouldBindJSON(&menu); err != nil {  
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  
		return  
	}  

	// Validasi input  
	if menu.Name == "" || menu.Price <= 0 || menu.ChefID <= 0 {  
		ctx.JSON(http.StatusBadRequest, gin.H{  
			"error": "Invalid menu details",  
		})  
		return  
	}  

	// Cek apakah chef tersedia  
	var chefExists int  
	err := c.db.QueryRow("SELECT COUNT(*) FROM chefs WHERE id = ?", menu.ChefID).Scan(&chefExists)  
	if err != nil || chefExists == 0 {  
		ctx.JSON(http.StatusBadRequest, gin.H{  
			"error": "Chef not found",  
		})  
		return  
	}  

	// Query insert  
	query := `INSERT INTO menus   
		(name, description, price, chef_id, category)   
		VALUES (?, ?, ?, ?, ?)`  
	  
	result, err := c.db.Exec(query,   
		menu.Name,   
		menu.Description,   
		menu.Price,   
		menu.ChefID,   
		menu.Category,  
	)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	// Dapatkan ID yang baru dibuat  
	id, _ := result.LastInsertId()  
	menu.ID = int(id)  

	ctx.JSON(http.StatusCreated, menu)  
}  

func (c *MenuController) GetAll(ctx *gin.Context) {  
	query := `SELECT id, name, description, price, chef_id, category FROM menus`  
	rows, err := c.db.Query(query)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  
	defer rows.Close()  

	var menus []models.Menu  
	for rows.Next() {  
		var menu models.Menu  
		err := rows.Scan(  
			&menu.ID,   
			&menu.Name,   
			&menu.Description,   
			&menu.Price,   
			&menu.ChefID,   
			&menu.Category,  
		)  
		if err != nil {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
			return  
		}  
		menus = append(menus, menu)  
	}  

	ctx.JSON(http.StatusOK, menus)  
}  

func (c *MenuController) GetByID(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	var menu models.Menu  
	query := `SELECT id, name, description, price, chef_id, category   
			  FROM menus WHERE id = ?`  
	  
	err := c.db.QueryRow(query, id).Scan(  
		&menu.ID,   
		&menu.Name,   
		&menu.Description,   
		&menu.Price,   
		&menu.ChefID,   
		&menu.Category,  
	)  
	if err != nil {  
		if err == sql.ErrNoRows {  
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})  
		} else {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		}  
		return  
	}  

	ctx.JSON(http.StatusOK, menu)  
}  

func (c *MenuController) Update(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	var menu models.Menu  
	if err := ctx.ShouldBindJSON(&menu); err != nil {  
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  
		return  
	}  

	// Validasi input  
	if menu.Name == "" || menu.Price <= 0 {  
		ctx.JSON(http.StatusBadRequest, gin.H{  
			"error": "Invalid menu details",  
		})  
		return  
	}  

	// Query update  
	query := `UPDATE menus   
			  SET name = ?, description = ?, price = ?, chef_id = ?, category = ?   
			  WHERE id = ?`  
	  
	_, err := c.db.Exec(query,   
		menu.Name,   
		menu.Description,   
		menu.Price,   
		menu.ChefID,   
		menu.Category,   
		id,  
	)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	// Set ID dari parameter  
	menu.ID, _ = strconv.Atoi(id)  

	ctx.JSON(http.StatusOK, menu)  
}  

func (c *MenuController) Delete(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	// Query delete  
	query := "DELETE FROM menus WHERE id = ?"  
	  
	result, err := c.db.Exec(query, id)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	// Periksa apakah ada baris yang terpengaruh  
	rowsAffected, err := result.RowsAffected()  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	if rowsAffected == 0 {  
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})  
		return  
	}  

	ctx.JSON(http.StatusOK, gin.H{  
		"message": "Menu deleted successfully",  
	})  
}  

func (c *MenuController) GetMenusByChef(ctx *gin.Context) {  
	chefID := ctx.Param("chefId")  
	  
	query := `SELECT id, name, description, price, chef_id, category   
			  FROM menus WHERE chef_id = ?`  
	  
	rows, err := c.db.Query(query, chefID)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  
	defer rows.Close()  

	var menus []models.Menu  
	for rows.Next() {  
		var menu models.Menu  
		err := rows.Scan(  
			&menu.ID,   
			&menu.Name,   
			&menu.Description,   
			&menu.Price,   
			&menu.ChefID,   
			&menu.Category,  
		)  
		if err != nil {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
			return  
		}  
		menus = append(menus, menu)  
	}  

	// Cek apakah daftar menu kosong  
	if len(menus) == 0 {  
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No menus found for this chef"})  
		return  
	}  

	ctx.JSON(http.StatusOK, menus)  
}  

func (c *MenuController) SearchMenus(ctx *gin.Context) {  
	// Parameter pencarian  
	name := ctx.Query("name")  
	category := ctx.Query("category")  
	minPrice := ctx.Query("min_price")  
	maxPrice := ctx.Query("max_price")  

	// Query dasar  
	query := `SELECT id, name, description, price, chef_id, category   
			  FROM menus WHERE 1=1`  
	  
	var args []interface{}  

	// Tambahkan kondisi pencarian  
	if name != "" {  
		query += " AND name LIKE ?"  
		args = append(args, "%"+name+"%")  
	}  

	if category != "" {  
		query += " AND category = ?"  
		args = append(args, category)  
	}  

	if minPrice != "" {  
		minPriceFloat, _ := strconv.ParseFloat(minPrice, 64)  
		query += " AND price >= ?"  
		args = append(args, minPriceFloat)  
	}  

	if maxPrice != "" {  
		maxPriceFloat, _ := strconv.ParseFloat(maxPrice, 64)  
		query += " AND price <= ?"  
		args = append(args, maxPriceFloat)  
	}  

	// Eksekusi query  
	rows, err := c.db.Query(query, args...)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  
	defer rows.Close()  

	var menus []models.Menu  
	for rows.Next() {  
		var menu models.Menu  
		err := rows.Scan(  
			&menu.ID,   
			&menu.Name,   
			&menu.Description,   
			&menu.Price,   
			&menu.ChefID,   
			&menu.Category,  
		)  
		if err != nil {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
			return  
		}  
		menus = append(menus, menu)  
	}  

	// Cek apakah daftar menu kosong  
	if len(menus) == 0 {  
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No menus found"})  
		return  
	}  

	ctx.JSON(http.StatusOK, menus)  
}