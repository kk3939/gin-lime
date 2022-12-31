package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/entity"
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

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo entity.Todo
	c.BindJSON(&todo)
	err := models.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo entity.Todo
	c.BindJSON(&todo)
	err := models.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo entity.Todo
	c.BindJSON(&todo)
	err := models.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}
