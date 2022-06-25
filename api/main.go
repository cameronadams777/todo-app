package main

import (
	"net/http"

	"github.com/cameronadams777/todo-app/controllers"
	"github.com/cameronadams777/todo-app/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		var response struct {
			Health string
		}
		response.Health = "I am healthy! 🚀"
		c.JSON(http.StatusOK, response)
	})

	router.GET("/todos", controllers.GetAllTodos)
	router.GET("/todos/:id", controllers.GetTodoById)
	router.POST("/todos", controllers.CreateNewTodo)
	router.POST("/todos/complete", controllers.CompleteTodo)

	router.Run(":5000")
}
