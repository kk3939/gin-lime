package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/controllers/todoController"
	"github.com/kk3939/gin-lime/controllers/userController"
	"github.com/kk3939/gin-lime/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		panic("JWT Error:" + err.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		ctx.JSON(404, gin.H{
			"message": "Not Found",
			"claims":  claims,
		})
	})

	r.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "signup",
		})
	})
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/logout", authMiddleware.LogoutHandler)
	// test
	h := r.Group("/hello").Use(authMiddleware.MiddlewareFunc())
	{
		h.GET("", middleware.HelloHandler)
	}

	// root
	r.GET("/", controllers.GetRoot)

	api := r.Group("/api/v1")
	{
		// todo
		api.GET("/todo", todoController.GetTodos)
		api.GET("/todo/:id", todoController.GetTodo)
		api.POST("/todo", todoController.CreateTodo)
		api.PUT("/todo/:id", todoController.UpdateTodo)
		api.DELETE("/todo/:id", todoController.DeleteTodo)
		// user
		api.GET("/user", userController.GetUsers)
		api.GET("/user/:id", userController.GetUser)
		api.POST("/user", userController.CreateUser)
		api.PUT("/user/:id", userController.UpdateUser)
		api.DELETE("/user/:id", userController.DeleteUser)
	}

	return r
}
