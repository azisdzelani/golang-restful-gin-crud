package handler

import (
	"fmt"
	"golang-restful-gin-crud/book"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Eko Kurniawan",
		"bio":  "software developer",
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"content":  "Hello Eko Kurniawan",
		"subtitle": "Belajar Golang Web API",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBooksHandler(ctx *gin.Context) {
	var booksInput book.BooksRequest

	err := ctx.ShouldBindJSON(&booksInput)

	if err != nil {

		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on failed %s, condition: %s", err.Field(), err.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": booksInput.Title,
		"price": booksInput.Price,
	})
}
