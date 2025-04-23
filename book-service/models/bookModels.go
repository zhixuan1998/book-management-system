package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BookBson struct {
	BookId      primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	Genre       string             `bson:"genre"`
	Description string             `bson:"description"`
	Isbn        string             `bson:"isbn"`
	Publisher   string             `bson:"publisher"`
	PublishedAt time.Time          `bson:"publishedAt"`
	IsDeleted   bool               `bson:"isDeleted"`
	ModifiedAt  time.Time          `bson:"modifiedAt"`
	ModifiedBy  primitive.ObjectID `bson:"modifiedBy"`
	CreatedAt   time.Time          `bson:"createdAt"`
	CreatedBy   primitive.ObjectID `bson:"createdBy"`
}

type BookJson struct {
	BookId      string `json:"bookId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	PublishedAt string `json:"publishedAt"`
	IsDeleted   bool   `json:"isDeleted"`
	ModifiedAt  string `json:"modifiedAt"`
	ModifiedBy  string `json:"modifiedBy"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
}

type BookBsonToAdd struct {
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	Genre       string             `bson:"genre"`
	Description string             `bson:"description"`
	Isbn        string             `bson:"isbn"`
	Publisher   string             `bson:"publisher"`
	PublishedAt time.Time          `bson:"publishedAt"`
	IsDeleted   bool               `bson:"isDeleted"`
	ModifiedAt  time.Time          `bson:"modifiedAt"`
	ModifiedBy  primitive.ObjectID `bson:"modifiedBy"`
	CreatedAt   time.Time          `bson:"createdAt"`
	CreatedBy   primitive.ObjectID `bson:"createdBy"`
}
