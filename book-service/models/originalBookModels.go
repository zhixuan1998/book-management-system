package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OriginalBookBson struct {
	OriginalBookId primitive.ObjectID `bson:"_id"`
	BookId         primitive.ObjectID `bson:"bookId"`
	Title          string             `bson:"title"`
	Author         string             `bson:"author"`
	Genre          string             `bson:"genre"`
	Description    string             `bson:"description"`
	Isbn           string             `bson:"isbn"`
	Publisher      string             `bson:"publisher"`
	PublishedAt    time.Time          `bson:"publishedAt"`
	IsDeleted      bool               `bson:"isDeleted"`
	CreatedAt      time.Time          `bson:"createdAt"`
	CreatedBy      primitive.ObjectID `bson:"createdBy"`
}

type OriginalBookJson struct {
	OriginalBookId string `json:"_id"`
	BookId         string `json:"bookId"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	Description    string `json:"description"`
	Isbn           string `json:"isbn"`
	Publisher      string `json:"publisher"`
	PublishedAt    string `json:"publishedAt"`
	IsDeleted      bool   `json:"isDeleted"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
}
