package main

import (
	"fmt"

	controller "Go_project/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	router.GET("/movies", controller.GetMovies())
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
