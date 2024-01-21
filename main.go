// main.go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Nonoyanping/mp3_to_text/handlers"
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
