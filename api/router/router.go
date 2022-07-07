package router

import (
	"api/controllers"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// Group all endpoints under /api path
	api := app.Group("/api")

	// Authentication
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	// Todos
	todos := api.Group("/todos", middleware.Protected())
	todos.GET("/", controllers.GetAllTodos)
	todos.GET("/:id", controllers.GetTodoById)
	todos.POST("/", controllers.CreateNewTodo)
	todos.POST("/complete", controllers.CompleteTodo)
}
