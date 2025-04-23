package book

import (
	"book-service/models"
	"time"
)

func convertTimeToISOString(timeInISOString time.Time) string {
	return timeInISOString.Format(time.RFC3339)
}

func ConvertToBookJson(originalBookBson models.OriginalBookBson) models.OriginalBookJson {
	return models.OriginalBookJson{
		OriginalBookId: originalBookBson.OriginalBookId.Hex(),
		BookId:         originalBookBson.BookId.Hex(),
		Title:          originalBookBson.Title,
		Author:         originalBookBson.Author,
		Genre:          originalBookBson.Genre,
		Description:    originalBookBson.Description,
		Isbn:           originalBookBson.Isbn,
		Publisher:      originalBookBson.Publisher,
		PublishedAt:    convertTimeToISOString(originalBookBson.PublishedAt),
		IsDeleted:      originalBookBson.IsDeleted,
		CreatedAt:      convertTimeToISOString(originalBookBson.CreatedAt),
		CreatedBy:      originalBookBson.CreatedBy.Hex(),
	}
}
