package controllers

import (
	"net/http"
	"time"

	"github.com/cameronadams777/todo-app/database"
	"github.com/cameronadams777/todo-app/models"
	"github.com/gin-gonic/gin"
)

// TODO: Paginate this request
func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	database.DB.First(&todo, id)
	c.JSON(http.StatusOK, todo)
}

func CreateNewTodo(c *gin.Context) {
	var createTodosInput struct {
		UserID      int    `json:"userId" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"-"`
	}

	c.BindJSON(&createTodosInput)

	todo := models.Todo{
		Title:       createTodosInput.Title,
		Description: createTodosInput.Description,
		UserID:      createTodosInput.UserID,
	}

	database.DB.Create(&todo)

	c.JSON(http.StatusCreated, todo)
}

func CompleteTodo(c *gin.Context) {
	var completeTodosInput struct {
		ID int `json:"id" binding:"required"`
	}

	c.BindJSON(&completeTodosInput)

	database.DB.Model(&models.Todo{}).Where("id = ?", completeTodosInput.ID).Update("completedAt", time.Now())

	c.JSON(http.StatusOK, gin.H{})
}
