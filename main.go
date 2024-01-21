// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"./handlers/"
)

func main() {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Load HTML templates from the static directory
	r.LoadHTMLGlob("static/*")

	// Define a route for the home page
	r.GET("/", handlers.IndexHandler)

	// Define a route for handling file uploads
	r.POST("/upload", handlers.UploadHandler)

	// Run the application on port 8080
	r.Run(":8080")
}
