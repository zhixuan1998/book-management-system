package book

import (
	"book-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getDefault() (primitive.ObjectID, time.Time) {
	admin, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	return admin, time.Now()
}

func convertISOStringToTime(timeInString string) time.Time {
	result, _ := time.Parse(time.RFC3339, timeInString)
	return result
}

func convertTimeToISOString(timeInISOString time.Time) string {
	return timeInISOString.Format(time.RFC3339)
}

func ConvertToBookBson(bookJson models.BookJson) models.BookBson {
	bookId, _ := primitive.ObjectIDFromHex(bookJson.BookId)
	modifiedBy, _ := primitive.ObjectIDFromHex(bookJson.ModifiedBy)
	createdBy, _ := primitive.ObjectIDFromHex(bookJson.CreatedBy)

	publishedAt := convertISOStringToTime(bookJson.PublishedAt)
	modifiedAt := convertISOStringToTime(bookJson.ModifiedAt)
	createdAt := convertISOStringToTime(bookJson.CreatedAt)

	return models.BookBson{
		BookId:      bookId,
		Title:       bookJson.Title,
		Author:      bookJson.Author,
		Genre:       bookJson.Genre,
		Description: bookJson.Description,
		Isbn:        bookJson.Isbn,
		Publisher:   bookJson.Publisher,
		PublishedAt: publishedAt,
		IsDeleted:   bookJson.IsDeleted,
		ModifiedAt:  modifiedAt,
		ModifiedBy:  modifiedBy,
		CreatedAt:   createdAt,
		CreatedBy:   createdBy,
	}
}

func ConvertToBookJson(bookBson models.BookBson) models.BookJson {

	return models.BookJson{
		BookId:      bookBson.BookId.Hex(),
		Title:       bookBson.Title,
		Author:      bookBson.Author,
		Genre:       bookBson.Genre,
		Description: bookBson.Description,
		Isbn:        bookBson.Isbn,
		Publisher:   bookBson.Publisher,
		PublishedAt: convertTimeToISOString(bookBson.PublishedAt),
		IsDeleted:   bookBson.IsDeleted,
		ModifiedAt:  convertTimeToISOString(bookBson.ModifiedAt),
		ModifiedBy:  bookBson.ModifiedBy.Hex(),
		CreatedAt:   convertTimeToISOString(bookBson.CreatedAt),
		CreatedBy:   bookBson.ModifiedBy.Hex(),
	}
}

func ConvertToBookBsonToAdd(bookJson models.BookJson) models.BookBsonToAdd {
	admin, date := getDefault()
	publishedAt := convertISOStringToTime(bookJson.PublishedAt)

	return models.BookBsonToAdd{
		Title:       bookJson.Title,
		Author:      bookJson.Author,
		Genre:       bookJson.Genre,
		Description: bookJson.Description,
		Isbn:        bookJson.Isbn,
		Publisher:   bookJson.Publisher,
		PublishedAt: publishedAt,
		ModifiedAt:  date,
		ModifiedBy:  admin,
		CreatedAt:   date,
		CreatedBy:   admin,
	}
}
