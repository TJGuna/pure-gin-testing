package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "1"},
	{ID: "2", Title: "The Lord of the Rings", Author: "2"},
	{ID: "3", Title: "The Wizard of Oz", Author: "3"},
}

func setupBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", getBooks)
		bookRoutes.GET("/:id", getBookByID) // New route for getting book by ID
		bookRoutes.POST("/", createBook)
		bookRoutes.PUT("/:id", updateBook)
		bookRoutes.DELETE("/:id", deleteBook)
	}
}
func getBookByID(c *gin.Context) {
	bookID := c.Param("id")

	for _, book := range books {
		if book.ID == bookID {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
	bookID := c.Param("id")

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == bookID {
			books[i] = updatedBook
			c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	bookID := c.Param("id")

	for i, book := range books {
		if book.ID == bookID {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
