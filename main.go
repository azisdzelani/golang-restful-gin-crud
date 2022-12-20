package main

import (
	"golang-restful-gin-crud/book"
	"golang-restful-gin-crud/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection errors")
	}

	db.AutoMigrate(&book.Book{})

	newRepository := book.NewRepository(db)
	newService := book.NewService(newRepository)
	newHandler := handler.NewBookHandler(newService)

	router := gin.Default()

	//api versioning
	v1 := router.Group("/v1")

	v1.GET("/books", newHandler.GetBooksHandler)
	v1.POST("/books", newHandler.CreateBookHandler)
	v1.GET("/books/:id", newHandler.GetBookByID)

	router.Run()
}
