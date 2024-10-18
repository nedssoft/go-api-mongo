package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
  ID        primitive.ObjectID `bson:"_id"`
	CreatedAt *time.Time `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
	Password string `bson:"password"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	
}

func (u *User) PreSave() {
	now := time.Now()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = &now
	u.UpdatedAt = &now
}
