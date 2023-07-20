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
	r.GET("/api/v1/todo", todoController.GetTodos)
	r.GET("/api/v1/todo/:id", todoController.GetTodo)
	r.POST("/api/v1/todo", todoController.CreateTodo)
	r.PUT("/api/v1/todo/:id", todoController.UpdateTodo)
	r.DELETE("/api/v1/todo/:id", todoController.DeleteTodo)
	// user
	r.GET("/api/v1/user", userController.GetUsers)
	r.GET("/api/v1/user/:id", userController.GetUser)
	r.POST("/api/v1/user", userController.CreateUser)
	r.PUT("/api/v1/user/:id", userController.UpdateUser)
	r.DELETE("/api/v1/user/:id", userController.DeleteUser)

	return r
}
