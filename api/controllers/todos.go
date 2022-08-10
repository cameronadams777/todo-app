package controllers

import (
	"api/database"
	"api/models"
	"api/structs"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: Paginate this request
func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	database.DB.Find(&todos, "user_id = ?", authScope.UserID)

	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	var todo models.Todo
	database.DB.First(&todo, id)

	if todo.UserID != authScope.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not own this todo item.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateNewTodo(c *gin.Context) {
	type CreateNewTodoInput struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"-"`
	}

	var input CreateNewTodoInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	todo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		UserID:      authScope.UserID,
	}

	database.DB.Create(&todo)

	c.JSON(http.StatusCreated, todo)
}

func CompleteTodo(c *gin.Context) {
	var completeTodosInput struct {
		ID int `json:"id" binding:"required"`
	}

	c.BindJSON(&completeTodosInput)

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	var todo models.Todo

	database.DB.First(&todo, completeTodosInput.ID)

	if todo.UserID != authScope.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not own this todo item.", "data": nil})
		return
	}

	database.DB.Model(&models.Todo{}).Where("id = ?", completeTodosInput.ID).Update("completedAt", time.Now())

	var completedTodo models.Todo

	database.DB.First(&todo, completeTodosInput.ID)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Todo item completed.", "data": completedTodo})
}
