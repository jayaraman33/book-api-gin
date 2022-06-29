package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{

	{ID: "1", Isbn: "123", Title: "Book one", Author: "Ram"},
	{ID: "2", Isbn: "456", Title: "Book two", Author: "Raja"},
	{ID: "3", Isbn: "789", Title: "Book three", Author: "vijay"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// get book

func getBook(c *gin.Context) {

	id := c.Param("id")

	for _, item := range books {

		if item.ID == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}

	}
}

// createBook

func createBook(c *gin.Context) {
	var newBook Book
	if error := c.BindJSON(&newBook); error != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

// deleteBook

func RemoveIndex(s []Book, index int) []Book {
	return append(s[:index], s[index+1:]...)
}

func deleteBook(c *gin.Context) {

	id := c.Param("id")

	for item, book := range books {
		if book.ID == id {
			books = RemoveIndex(books, item)
			return

		}
	}
}

// updateBook

func updateBook(c *gin.Context) {

	id := c.Param("id")

	var updateData Book

	if error := c.BindJSON(&updateData); error != nil {

		return

	}

	for item, book := range books {
		if book.ID == id {
			books[item] = updateData
			c.IndentedJSON(http.StatusOK, books[item])
			return
		}
	}

}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		r.GET("/books", getBooks)

		r.GET("/books/:id", getBook)

		r.POST("/books/", createBook)

		r.DELETE("/books/:id", deleteBook)

		r.PUT("/books/:id", updateBook)

	})
	r.Run()
}
