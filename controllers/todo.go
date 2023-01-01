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
	todo := entity.Todo{}
	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	if todoGot, err := models.GetTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	} else {
		todo := entity.Todo{}
		c.BindJSON(&todo)
		todo.Id = todoGot.Id
		if err := models.UpdateTodo(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	if todoGot, err := models.GetTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmtErrMsg(err),
		})
		return
	} else {
		todo := entity.Todo{}
		todo.Id = todoGot.Id
		if err := models.DeleteTodo(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}
