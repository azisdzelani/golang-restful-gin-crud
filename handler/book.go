package handler

import (
	"fmt"
	"golang-restful-gin-crud/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooksHandler(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	bookData, err := h.bookService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	bookResponse := book.BookResponse{
		ID:          bookData.ID,
		Title:       bookData.Title,
		Description: bookData.Description,
		Price:       bookData.Price,
		Rating:      bookData.Rating,
		Discount:    bookData.Discount,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBooksHandler(ctx *gin.Context) {

}

func (h *bookHandler) CreateBookHandler(ctx *gin.Context) {
	var booksRequest book.BooksRequest

	err := ctx.ShouldBindJSON(&booksRequest)

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

	book, err := h.bookService.Create(booksRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
