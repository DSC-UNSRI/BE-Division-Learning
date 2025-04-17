package controllers  

import (  
	"database/sql"  
	"fmt"  
	"net/http"  
	"time"  

	"restaurant-backend/models"  
	"restaurant-backend/utils"  

	"github.com/gin-gonic/gin"  
)  

type ChefController struct {  
	db *sql.DB  
}  

func NewChefController(db *sql.DB) *ChefController {  
	return &ChefController{db: db}  
}  

func (c *ChefController) Create(ctx *gin.Context) {  
	var chef models.Chef  
	if err := ctx.ShouldBindJSON(&chef); err != nil {  
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  
		return  
	}  

	// Validasi input  
	if chef.Name == "" || chef.Username == "" || chef.PasswordHash == "" {  
		ctx.JSON(http.StatusBadRequest, gin.H{  
			"error": "Name, username, and password are required",  
		})  
		return  
	}  

	// Cek apakah username sudah ada  
	var count int  
	err := c.db.QueryRow("SELECT COUNT(*) FROM chefs WHERE username = ?", chef.Username).Scan(&count)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  
	if count > 0 {  
		ctx.JSON(http.StatusConflict, gin.H{  
			"error": "Username already exists",  
		})  
		return  
	}  

	// Query insert  
	query := `INSERT INTO chefs   
		(name, speciality, experience, username, password_hash)   
		VALUES (?, ?, ?, ?, ?)`  
	  
	result, err := c.db.Exec(query,   
		chef.Name,   
		chef.Speciality,   
		chef.Experience,   
		chef.Username,   
		chef.PasswordHash,  
	)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	// Dapatkan ID yang baru dibuat  
	id, _ := result.LastInsertId()  
	chef.ID = int(id)  

	// Hapus password sebelum mengirim response  
	chef.PasswordHash = ""  

	ctx.JSON(http.StatusCreated, chef)  
}  

func (c *ChefController) GetAll(ctx *gin.Context) {  
	query := "SELECT id, name, speciality, experience, username FROM chefs"  
	rows, err := c.db.Query(query)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  
	defer rows.Close()  

	var chefs []models.Chef  
	for rows.Next() {  
		var chef models.Chef  
		err := rows.Scan(  
			&chef.ID,   
			&chef.Name,   
			&chef.Speciality,   
			&chef.Experience,   
			&chef.Username,  
		)  
		if err != nil {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
			return  
		}  
		chefs = append(chefs, chef)  
	}  

	ctx.JSON(http.StatusOK, chefs)  
}  

func (c *ChefController) GetByID(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	var chef models.Chef  
	query := `SELECT id, name, speciality, experience, username   
			  FROM chefs WHERE id = ?`  
	  
	err := c.db.QueryRow(query, id).Scan(  
		&chef.ID,   
		&chef.Name,   
		&chef.Speciality,   
		&chef.Experience,   
		&chef.Username,  
	)  
	if err != nil {  
		if err == sql.ErrNoRows {  
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Chef not found"})  
		} else {  
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		}  
		return  
	}  

	ctx.JSON(http.StatusOK, chef)  
}  

func (c *ChefController) Update(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	var chef models.Chef  
	if err := ctx.ShouldBindJSON(&chef); err != nil {  
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  
		return  
	}  

	// Query update  
	query := `UPDATE chefs   
			  SET name = ?, speciality = ?, experience = ?   
			  WHERE id = ?`  
	  
	_, err := c.db.Exec(query,   
		chef.Name,   
		chef.Speciality,   
		chef.Experience,   
		id,  
	)  
	if err != nil {  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})  
		return  
	}  

	// Set ID dari parameter  
	chef.ID, _ = strconv.Atoi(id)  

	ctx.JSON(http.StatusOK, chef)  
}  

func (c *ChefController) Delete(ctx *gin.Context) {  
	id := ctx.Param("id")  
	  
	// Query delete  
	query := "DELETE FROM chefs WHERE id = ?"  
	  
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Chef not found"})  
		return  
	}  

	ctx.JSON(http.StatusOK, gin.H{  
		"message": "Chef deleted successfully",  
	})  
}  

func (c *ChefController) Login(ctx *gin.Context) {  
	var loginRequest struct {  
		Username string `json:"username"`  
		Password string `json:"password"`  
	}  

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {  
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})  
		return  
	}  

	// Cek kredensial  
	var chef models.Chef  
	query := `SELECT id, name, username, speciality, experience, password_hash   
			  FROM chefs WHERE username = ? AND password_hash = ?`  
	  
	err := c.db.QueryRow(query, loginRequest.Username, loginRequest.Password).Scan(  
		&chef.ID,   
		&chef.Name,   
		&chef.Username,   
		&chef.Speciality,   
		&chef.Experience,  
		&chef.PasswordHash,  
	)  
	if err != nil {  
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})  
		return  
	}  

	// Hapus password sebelum mengirim response  
	chef.PasswordHash = ""  

	ctx.JSON(http.StatusOK, chef)  
}