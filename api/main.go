package main

import (
	"api/database"
	"api/router"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	app := gin.Default()

	app.Use(cors.Default()) // TODO: Configure this later to lock it down better

	app.GET("/health", func(c *gin.Context) {
		var response struct {
			Health string
		}
		response.Health = "I am healthy! 🚀"
		c.JSON(http.StatusOK, response)
	})

	router.SetupRouter(app)

	app.Run(":5000")
}
