package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db *DB, ctx context.Context) {
	CreatePostIndexes(db, ctx)
	CreateUserIndexes(db, ctx)
}

func CreatePostIndexes(db *DB, ctx context.Context) {
	collection := db.GetCollection("posts")
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"created_at", 1}, {"user_id", 1}},
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Failed to create index: %v", err)
	}
}

func CreateUserIndexes(db *DB, ctx context.Context) {
	collection := db.GetCollection("users")
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"email", 1}},
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Failed to create index: %v", err)
	}
}
