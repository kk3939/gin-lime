package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/models"
)

func GetTodos(c *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}
