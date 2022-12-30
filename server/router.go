package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetRoot)
	r.GET("/todo", controllers.GetTodos)
	return r
}
