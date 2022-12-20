package book

import "encoding/json"

type BooksRequest struct {
	// directive json => variable SubTitle dipakai u/
	// menangkap json sub_title
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"required,number"`
	Discount    int         `json:"discount" binding:"required,number"`
}
