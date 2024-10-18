package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID  primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Body string `bson:"body"`
	UserID  primitive.ObjectID `bson:"user_id"`
	CreatedAt *time.Time `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

func (p *Post) PreSave() {
	now := time.Now()
	p.ID = primitive.NewObjectID()
	p.CreatedAt = &now
	p.UpdatedAt = &now
}
