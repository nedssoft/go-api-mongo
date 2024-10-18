package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostResponse struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	UserID string `bson:"user_id" json:"user_id"`
	CreatedAt *time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

type PostsResponse struct {
	Posts []PostResponse `json:"posts"`
}

type PostUserResponse struct {
	PostResponse
	User UserResponse `json:"user"`
}

