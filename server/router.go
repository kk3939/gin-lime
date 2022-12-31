package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetRoot)

	r.GET("/todo", controllers.GetTodos)
	r.GET("/todo/:id", controllers.GetTodo)
	r.POST("/todo/:id", controllers.CreateTodo)
	r.PUT("/todo/:id", controllers.UpdateTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)
	return r
}
