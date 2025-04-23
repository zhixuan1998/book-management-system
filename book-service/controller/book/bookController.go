package book

import (
	"book-service/models"
	"book-service/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	repositories *models.RepositoryContainer
}

func CreateBookController(repositories *models.RepositoryContainer) *BookController {
	return &BookController{repositories: repositories}
}

func (controller *BookController) CreateBook(c *gin.Context) {
	var bookJson models.BookJson

	if error := c.ShouldBindJSON(&bookJson); error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	result, error := controller.repositories.BookRepository.CreateBook(c, bookJson)

	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	response := utils.GenerateSuccessResponseWithData(result)

	c.JSON(http.StatusOK, response)
}

func (controller *BookController) GetBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip := (page - 1) * limit

	books, error := controller.repositories.BookRepository.GetBooks(c, limit+1, skip)

	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	count := len(books)

	if count > limit {
		books = books[:count-1]
	}

	pagination := utils.GeneratePagination(page, limit, count)

	response := utils.GenerateSuccessResponseWithPagination(books, pagination)

	c.JSON(http.StatusOK, response)
}

func (controller *BookController) GetBook(c *gin.Context) {
	bookId := c.Param("bookId")

	book, error := controller.repositories.BookRepository.GetBook(c, bookId)

	if error != nil {
		fmt.Println(error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	if book == nil {
		fmt.Println(error)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateErrorResponse())
		return
	}

	response := utils.GenerateSuccessResponseWithData(book)

	c.JSON(http.StatusOK, response)
}

func (controller *BookController) UpdateBook(c *gin.Context) {
	bookId := c.Param("bookId")

	var bookJson models.BookJson
	bookJson.BookId = bookId

	if error := c.ShouldBindJSON(&bookJson); error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	book, bookError := controller.repositories.BookRepository.GetBook(c, bookId)

	if bookError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	if book == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateErrorResponse())
		return
	}

	_, updateError := controller.repositories.BookRepository.UpdateBook(c, bookJson)

	if updateError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	response := utils.GenerateSuccessResponse()

	c.JSON(http.StatusOK, response)
}

func (controller *BookController) DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")

	book, bookError := controller.repositories.BookRepository.GetBook(c, bookId)

	if bookError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	if book == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.GenerateErrorResponse())
		return
	}

	_, deleteError := controller.repositories.BookRepository.DeleteBook(c, bookId)

	if deleteError != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	response := utils.GenerateSuccessResponse()

	c.JSON(http.StatusOK, response)
}

func (controller *BookController) ResetBookList(c *gin.Context) {
	originalBooks, error := controller.repositories.OriginalBookRepository.GetOriginalBooks(c)

	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenerateErrorResponse())
		return
	}

	if len(originalBooks) != 0 {
		bookIdsToBeRemained := []string{}

		for _, originalBook := range originalBooks {
			book := models.BookJson{
				BookId:      originalBook.BookId,
				Title:       originalBook.Title,
				Author:      originalBook.Author,
				Genre:       originalBook.Genre,
				Description: originalBook.Description,
				Isbn:        originalBook.Isbn,
				Publisher:   originalBook.Publisher,
				PublishedAt: originalBook.PublishedAt,
			}

			controller.repositories.BookRepository.UpdateBook(c, book)
			bookIdsToBeRemained = append(bookIdsToBeRemained, originalBook.BookId)
		}

		controller.repositories.BookRepository.DeleteBooks(c, bookIdsToBeRemained)

	}

	response := utils.GenerateSuccessResponse()
	c.JSON(http.StatusOK, response)
}
