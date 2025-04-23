package repositories

import (
	bookController "book-service/controller/book"
	"book-service/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookRepo struct {
	books *mongo.Collection
}

func CreateBookRepository(db *mongo.Database) *BookRepo {
	return &BookRepo{
		books: db.Collection("books"),
	}
}

func (r *BookRepo) CreateBook(ctx context.Context, bookJson models.BookJson) (*models.BookJson, any) {
	bookBson := bookController.ConvertToBookBsonToAdd(bookJson)

	insertedResult, insertResultError := r.books.InsertOne(ctx, bookBson)

	if insertResultError != nil {
		return nil, insertResultError
	}

	insertedId := insertedResult.InsertedID

	var result models.BookBson

	getResultError := r.books.FindOne(ctx, bson.M{"_id": insertedId}).Decode(&result)

	if getResultError != nil {
		return nil, getResultError
	}

	resultInBookJson := bookController.ConvertToBookJson(result)

	return &resultInBookJson, nil
}

func (r *BookRepo) GetBooks(ctx context.Context, limit, skip int) ([]models.BookJson, any) {

	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	fmt.Println(int64(limit), int64(skip))

	cursor, error := r.books.Find(ctx, bson.M{"isDeleted": false}, opts)

	if error != nil {
		return []models.BookJson{}, error
	}

	defer cursor.Close(ctx)

	resultsInBson := []models.BookBson{}

	if error = cursor.All(context.TODO(), &resultsInBson); error != nil {
		return nil, error
	}

	resultsInJson := []models.BookJson{}

	for _, bookBson := range resultsInBson {
		resultsInJson = append(resultsInJson, bookController.ConvertToBookJson(bookBson))
	}

	return resultsInJson, nil
}

func (r *BookRepo) GetBook(ctx context.Context, bookId string) (*models.BookJson, any) {
	bookIdInObjectId, _ := primitive.ObjectIDFromHex(bookId)

	var result models.BookBson

	error := r.books.FindOne(ctx, bson.M{"_id": bookIdInObjectId, "isDeleted": false}).Decode(&result)

	if error != nil {
		return nil, error
	}

	resultInJson := bookController.ConvertToBookJson(result)

	return &resultInJson, nil
}

func (r *BookRepo) UpdateBook(ctx context.Context, bookJson models.BookJson) (bool, any) {
	bookId, _ := primitive.ObjectIDFromHex(bookJson.BookId)
	bookBson := bookController.ConvertToBookBson(bookJson)

	filterQuery := bson.M{
		"_id":       bookId,
		"isDeleted": false,
	}

	admin, date := getDefaultUpdateField()

	updateQuery := bson.M{
		"$set": bson.M{
			"title":       bookBson.Title,
			"author":      bookBson.Author,
			"genre":       bookBson.Genre,
			"description": bookBson.Description,
			"isbn":        bookBson.Isbn,
			"publisher":   bookBson.Publisher,
			"publishedAt": bookBson.PublishedAt,
			"modifiedBy":  admin,
			"modifiedAt":  date,
		},
	}

	_, error := r.books.UpdateOne(ctx, filterQuery, updateQuery)

	if error != nil {
		return false, error
	}

	return true, nil
}

func (r *BookRepo) DeleteBook(ctx context.Context, bookId string) (bool, any) {
	bookIdInObjectId, _ := primitive.ObjectIDFromHex(bookId)

	filterQuery := bson.M{"_id": bookIdInObjectId}

	updateQuery := bson.M{
		"$set": bson.M{
			"isDeleted": true,
		},
	}

	_, error := r.books.UpdateOne(ctx, filterQuery, updateQuery)

	if error != nil {
		return false, error
	}

	return true, nil
}

func (r *BookRepo) DeleteBooks(ctx context.Context, excludedBookIds []string) (bool, any) {
	bookObjectIds := []primitive.ObjectID{}

	for _, bookId := range excludedBookIds {
		objectId, _ := primitive.ObjectIDFromHex(bookId)
		bookObjectIds = append(bookObjectIds, objectId)
	}

	filterQuery := bson.M{
		"_id": bson.M{
			"$nin": bookObjectIds,
		},
	}

	admin, date := getDefaultUpdateField()

	updateQuery := bson.M{
		"$set": bson.M{
			"isDeleted":  true,
			"modifiedBy": admin,
			"modifiedAt": date,
		},
	}

	_, error := r.books.UpdateMany(ctx, filterQuery, updateQuery)

	if error != nil {
		return false, error
	}

	return true, nil
}

func getDefaultUpdateField() (primitive.ObjectID, time.Time) {
	admin, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	return admin, time.Now()
}
