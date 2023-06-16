package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route and its handler function
	router.GET("/", helloHandler)

	// Start the server
	router.Run(":8080")
}

func helloHandler(c *gin.Context) {
	// Set the response status and data
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
