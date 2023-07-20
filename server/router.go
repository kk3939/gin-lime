package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/controllers/todoController"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetRoot)

	r.GET("/todo", todoController.GetTodos)
	r.GET("/todo/:id", todoController.GetTodo)
	r.POST("/todo", todoController.CreateTodo)
	r.PUT("/todo/:id", todoController.UpdateTodo)
	r.DELETE("/todo/:id", todoController.DeleteTodo)
	return r
}
