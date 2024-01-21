// handlers/upload_handler.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"mp3-to-text-webapp/services"
	"net/http"
)

// UploadHandler handles file uploads
func UploadHandler(c *gin.Context) {
	// Retrieve the uploaded file from the form data
	file, err := c.FormFile("audio")
	if err != nil {
		// Handle bad requests with an error message
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"error": err.Error()})
		return
	}

	// Save the uploaded file to the 'uploads' directory
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		// Handle internal server errors with an error message
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": err.Error()})
		return
	}

	// Convert MP3 to text using a service function
	textResult := services.ConvertMP3ToText("uploads/" + file.Filename)

	// Render the index.html template with success message and text result
	c.HTML(http.StatusOK, "index.html", gin.H{"message": "File uploaded successfully", "textResult": textResult})
}
