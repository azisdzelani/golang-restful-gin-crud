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

	bookRequest := book.BooksRequest{
		Title: "Buku Gundam",
		Price: 200000,
	}

	newService.Create(bookRequest)

	//CRUD

	// ============
	// CREATE data
	// ============
	//
	// book := book.Book{}
	// book.Title = "Atomic Habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Buku ini tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("===========================")
	// }

	// ============
	// READ data
	// ============
	//
	// var books []book.Book
	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("===========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Printf("Book Object : %v", b)
	// }

	// ============
	// UPDATE data
	// ============
	//
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("===========================")
	// }

	// book.Title = "Man Tiger (revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error Updating book record")
	// 	fmt.Println("===========================")
	// }

	// ============
	// DELETE data
	// ============
	//
	// var book book.Book

	// err = db.Debug().Where("id", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error Finding book record")
	// 	fmt.Println("===========================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("===========================")
	// 	fmt.Println("Error Delete book record")
	// 	fmt.Println("===========================")
	// }

	router := gin.Default()

	//api versioning
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
