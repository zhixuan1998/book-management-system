package models

import (
	"context"
)

type BookRepository interface {
	CreateBook(ctx context.Context, bookJson BookJson) (*BookJson, any)
	GetBook(ctx context.Context, bookId string) (*BookJson, any)
	GetBooks(ctx context.Context, limit, skip int) ([]BookJson, any)
	UpdateBook(ctx context.Context, bookJson BookJson) (bool, any)
	DeleteBook(ctx context.Context, bookId string) (bool, any)
	DeleteBooks(ctx context.Context, excludedBookIds []string) (bool, any)
}

type OriginalBookRepository interface {
	GetOriginalBooks(ctx context.Context) ([]OriginalBookJson, any)
}
