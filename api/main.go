package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	CompletedAt time.Time
	UserID      int
}

func main() {
	dsn := "host=localhost user=postgres password=hzKD2b#KXk!hhseB=+U7P?A? dbname=todoapp_db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Todo{})

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		var response struct {
			Health string
		}
		response.Health = "Fully Healthy!"
		c.JSON(http.StatusOK, response)
	})

	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	router.GET("/todo/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo
		db.First(&todo, id)
		c.JSON(http.StatusOK, todo)
	})

	router.POST("/todo", func(c *gin.Context) {
		var createTodoInput struct {
			UserID      int    `json:"userId" binding:"required"`
			Title       string `json:"title" binding:"required"`
			Description string `json:"description" binding:"-"`
		}

		c.BindJSON(&createTodoInput)

		todo := Todo{
			Title:       createTodoInput.Title,
			Description: createTodoInput.Description,
			UserID:      createTodoInput.UserID,
		}

		db.Create(&todo)

		c.JSON(http.StatusOK, todo)
	})

	router.Run(":5000")
}
