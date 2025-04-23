package main

import (
	"book-service/controller/book"
	"book-service/database"
	"book-service/models"
	"book-service/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	mongoURI := "mongodb+srv://dummyuser:GDiwWGETEV3zlIe0@dummy-use.lh6wmdp.mongodb.net"
	db := database.Connect(mongoURI)

	repositoryContainer := &models.RepositoryContainer{
		BookRepository: repositories.CreateBookRepository(db),
		OriginalBookRepository: repositories.CreateOriginalBookRepository(db),
	}

	bookController := book.CreateBookController(repositoryContainer)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}))

	r := router.Group("/api/v1")
	r.POST("/books", bookController.CreateBook)
	r.GET("/books", bookController.GetBooks)
	r.GET("/books/:bookId", bookController.GetBook)
	r.PUT("/books/:bookId", bookController.UpdateBook)
	r.DELETE("/books/:bookId", bookController.DeleteBook)
	r.POST("/books/reset", bookController.ResetBookList)

	router.Run(":8080")
}
