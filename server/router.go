package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/controllers/todoController"
	"github.com/kk3939/gin-lime/controllers/userController"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	// root
	r.GET("/", controllers.GetRoot)
	// todo
	r.GET("/todo", todoController.GetTodos)
	r.GET("/todo/:id", todoController.GetTodo)
	r.POST("/todo", todoController.CreateTodo)
	r.PUT("/todo/:id", todoController.UpdateTodo)
	r.DELETE("/todo/:id", todoController.DeleteTodo)
	// user
	r.GET("/user", userController.GetUsers)
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	return r
}
