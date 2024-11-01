package service

import (
	"context"
	"errors"
	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/data/requests"
	"github.com/nedssoft/go-api-mongo/data/responses"
	"github.com/nedssoft/go-api-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(db *db.DB, ctx context.Context) *UserService {
	userCollection := db.GetCollection("users")
	return &UserService{userCollection: userCollection, ctx: ctx}
}

func (s *UserService) CreateUser(payload *requests.UserPayload) (userResponse *responses.UserResponse, err error) {
	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	user.PreSave()
	result, err := s.userCollection.InsertOne(s.ctx, user)
	if err != nil {
		return nil, err
	}
	return &responses.UserResponse{
		ID:        result.InsertedID.(primitive.ObjectID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *UserService) GetUserWithPosts(id primitive.ObjectID) (*responses.UserPostsResponse, error) {
	res, err := s.userCollection.Aggregate(s.ctx, mongo.Pipeline{
		{{"$match", bson.D{{"_id", id}}}},
		{{"$lookup", bson.M{
			"from": "posts",
			"localField": "_id",
			"foreignField": "user_id",
			"as": "posts",
		}}},
		{{"$project", bson.D{
			{"user", bson.M{
				"_id": "$_id",
				"name": "$name",
				"email": "$email",
				"created_at": "$created_at",
				"updated_at": "$updated_at",
			}},
			{"posts", "$posts"},
		}}},
	})
	if err != nil {
		return nil, err
	}
	defer res.Close(s.ctx)
	var result responses.UserPostsResponse
	if res.Next(s.ctx) {
		if err := res.Decode(&result); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("user not found")
	}
	return &result, nil
}

func (s *UserService) GetUsers() ([]responses.UserResponse, error) {
	var users []responses.UserResponse
	cursor, err := s.userCollection.Find(s.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)
	for cursor.Next(s.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return users, nil
}

func (s *UserService) DeleteUser(id primitive.ObjectID) error {
	_, err := s.userCollection.DeleteOne(s.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(id primitive.ObjectID, user *requests.UserUpdatePayload) error {
	_, err := s.userCollection.UpdateOne(s.ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.userCollection.FindOne(s.ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserById(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	if err := s.userCollection.FindOne(s.ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}


