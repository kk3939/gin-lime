package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is root endpoint.",
	})
}
