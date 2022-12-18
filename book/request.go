package book

type BooksRequest struct {
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
	//SubTitle string `json:"sub_title"` //directive json => variable SubTitle dipakai u/ menangkap json sub_title
}
