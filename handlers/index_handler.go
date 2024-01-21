// handlers/index_handler.go
package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler handles requests for the home page
func IndexHandler(c *gin.Context) {
	// Render the index.html template
	c.HTML(http.StatusOK, "index.html", nil)
}
