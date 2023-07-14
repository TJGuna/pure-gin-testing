package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Enable CORS middleware
	r.Use(enableCors)

	// Routes
	setupBookRoutes(r)
	setupAuthorRoutes(r)

	r.Run(":8080")
}

// enableCors enables CORS (Cross-Origin Resource Sharing) for the server.
func enableCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")

	// Handle OPTIONS request
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
