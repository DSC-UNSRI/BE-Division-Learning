package controllers

import (
	"net/http"
	"playlist-app/models"
	"playlist-app/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func (ctrl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.Repo.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.Repo.FindByID(parseUint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.Repo.Delete(parseUint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *UserController) Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.Repo.FindByName(input.Name)
	if err != nil || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}

// Helper
func parseUint(s string) uint {
	n, _ := strconv.Atoi(s)
	return uint(n)
}
