package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/controllers/tdc"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetRoot)

	r.GET("/todo", tdc.GetTodos)
	r.GET("/todo/:id", tdc.GetTodo)
	r.POST("/todo", tdc.CreateTodo)
	r.PUT("/todo/:id", tdc.UpdateTodo)
	r.DELETE("/todo/:id", tdc.DeleteTodo)
	return r
}
