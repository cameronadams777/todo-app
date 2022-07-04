package router

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	api := app.Group("/api")

	// Todos
	todos := api.Group("/todos")
	todos.GET("/todos", controllers.GetAllTodos)
	todos.GET("/todos/:id", controllers.GetTodoById)
	todos.POST("/todos", controllers.CreateNewTodo)
	todos.POST("/todos/complete", controllers.CompleteTodo)
}
