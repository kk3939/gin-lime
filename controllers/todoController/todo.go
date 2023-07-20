package todoController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/models/todoModel"
)

func GetTodos(c *gin.Context) {
	todos, err := todoModel.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := todoModel.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	todo := entity.Todo{}
	c.BindJSON(&todo)
	err := todoModel.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	if todoGot, err := todoModel.GetTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	} else {
		todo := entity.Todo{}
		c.BindJSON(&todo)
		todo.Id = todoGot.Id
		if err := todoModel.UpdateTodo(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": controllers.FmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	if todoGot, err := todoModel.GetTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	} else {
		todo := entity.Todo{}
		todo.Id = todoGot.Id
		if err := todoModel.DeleteTodo(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": controllers.FmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}
