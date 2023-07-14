package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var authors = []Author{
	{ID: "1", Name: "Author 1"},
	{ID: "2", Name: "Author 2"},
	{ID: "3", Name: "Author 3"},
}

func setupAuthorRoutes(r *gin.Engine) {
	authorRoutes := r.Group("/authors")
	{
		authorRoutes.GET("/", getAuthors)
		authorRoutes.GET("/:id", getAuthorByID) // New route for getting author by ID
		authorRoutes.POST("/", createAuthor)
		authorRoutes.PUT("/:id", updateAuthor)
		authorRoutes.DELETE("/:id", deleteAuthor)
	}
}

func getAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}
func getAuthorByID(c *gin.Context) {
	authorID := c.Param("id")

	for _, author := range authors {
		if author.ID == authorID {
			c.JSON(http.StatusOK, author)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}
func createAuthor(c *gin.Context) {
	var author Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authors = append(authors, author)
	c.JSON(http.StatusCreated, author)
}

func updateAuthor(c *gin.Context) {
	authorID := c.Param("id")

	var updatedAuthor Author
	if err := c.ShouldBindJSON(&updatedAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, author := range authors {
		if author.ID == authorID {
			authors[i] = updatedAuthor
			c.JSON(http.StatusOK, gin.H{"message": "Author updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}

func deleteAuthor(c *gin.Context) {
	authorID := c.Param("id")

	for i, author := range authors {
		if author.ID == authorID {
			authors = append(authors[:i], authors[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
}
