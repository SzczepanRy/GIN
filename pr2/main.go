package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `josn:"title"`
	Author   string `josn:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// ctx is like teh req and res in js
func getBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func bookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := getBookById(id)

	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("id not found")
}

func checkoutBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing param"})
	}

	book, err := getBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not availible"})
		return
	}

	book.Quantity -= 1

	ctx.IndentedJSON(http.StatusOK, book)

}

func returnBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")

	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing param"})
	}

	book, err := getBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1

	ctx.IndentedJSON(http.StatusOK, book)

}

func createBook(ctx *gin.Context) {
	var newBook book

	//binds the body of req to the newBook value
	if err := ctx.BindJSON(&newBook); err != nil {
		//returns static err
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	books = append(books, newBook)

	ctx.IndentedJSON(http.StatusCreated, newBook)

}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/createBook", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:3000")

}
