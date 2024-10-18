package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt *time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

type UserPostsResponse struct {
	User UserResponse `json:"user"`
	Posts []PostResponse `json:"posts"`
}
