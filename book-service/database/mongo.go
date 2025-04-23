package database

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, error := mongo.Connect(ctx, clientOptions)

	if error != nil {
		log.Fatal("MongoDB connection error:", error)
	}

	if error = client.Ping(ctx, nil); error != nil {
		log.Fatal("MongoDB ping failed", error)
	}

	return client.Database("sample")
}