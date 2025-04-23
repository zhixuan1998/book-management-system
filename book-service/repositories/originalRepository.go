package repositories

import (
	originalBookController "book-service/controller/originalBook"
	"book-service/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OriginalBookRepo struct {
	collection *mongo.Collection
}

func CreateOriginalBookRepository(db *mongo.Database) *OriginalBookRepo {
	return &OriginalBookRepo{collection: db.Collection("originalBooks")}
}

func (r *OriginalBookRepo) GetOriginalBooks(ctx context.Context) ([]models.OriginalBookJson, any) {
	cursor, error := r.collection.Find(ctx, bson.M{})

	if error != nil {
		return []models.OriginalBookJson{}, error
	}

	defer cursor.Close(ctx)

	resultsInBson := []models.OriginalBookBson{}

	if error := cursor.All(ctx, &resultsInBson); error != nil {
		return []models.OriginalBookJson{}, error
	}

	resultsInJson := []models.OriginalBookJson{}

	for _, originalBookBson := range resultsInBson {
		resultsInJson = append(resultsInJson, originalBookController.ConvertToBookJson(originalBookBson))
	}

	return resultsInJson, nil
}
