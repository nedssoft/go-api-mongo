package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
	Db *mongo.Database
}

func NewDB(mongoURI string, dbName string) (*DB, error) {
	clinetOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clinetOptions)
	if err != nil {
		return nil, err
	}
	db := &DB{
		Client: client,
		Db: client.Database(dbName),
	}
	return db, nil
}

func (d *DB) GetCollection (collectionName string) *mongo.Collection {
	return d.Db.Collection(collectionName)
}
