package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// HandleError - Menangani error umum dan mengirim response standar
func HandleError(ctx *gin.Context, statusCode int, errMessage string) {
	log.Println("Error:", errMessage)
	ctx.JSON(statusCode, gin.H{
		"error": errMessage,
	})
}
