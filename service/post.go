package service

import (
	"context"
	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/data/requests"
	"github.com/nedssoft/go-api-mongo/data/responses"
	"github.com/nedssoft/go-api-mongo/models"
	"github.com/nedssoft/go-api-mongo/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	postCollection *mongo.Collection
	ctx context.Context
}

func NewPostService(db *db.DB, ctx context.Context) *PostService {
	postCollection := db.GetCollection("posts")
	return &PostService{postCollection: postCollection, ctx: ctx}
}

func (s *PostService) CreatePost(payload *requests.PostPayload,userId primitive.ObjectID) (postResponse *responses.PostResponse, err error) {
	now := time.Now()
	post := models.Post{
		Title:  payload.Title,
		Body:   payload.Body,
		UserID: userId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	post.PreSave()
	result, err := s.postCollection.InsertOne(s.ctx, post)
	if err != nil {
		return nil, err
	}
	return &responses.PostResponse{
		ID: result.InsertedID.(primitive.ObjectID),
		Title: post.Title,
		Body: post.Body,
		UserID: post.UserID.Hex(),
		CreatedAt: utils.DefaultValue(post.CreatedAt, &now),
		UpdatedAt: utils.DefaultValue(post.UpdatedAt, &now),
	}, nil
}

func (s *PostService) GetPost(id primitive.ObjectID) (*responses.PostResponse, error) {
	var post responses.PostResponse
	if err := s.postCollection.FindOne(s.ctx, bson.M{"_id": id}).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) GetPosts() ([]responses.PostResponse, error) {
	var posts []responses.PostResponse
	cursor, err := s.postCollection.Find(s.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)
	for cursor.Next(s.ctx) {
		var post responses.PostResponse
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *PostService) DeletePost(id primitive.ObjectID) error {
	_, err := s.postCollection.DeleteOne(s.ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) UpdatePost(id primitive.ObjectID, payload *requests.PostUpdatePayload, userId primitive.ObjectID) error {
	post := bson.M{
		"title":  payload.Title,
		"body":   payload.Body,
		"updated_at": time.Now(),
	}
	result := s.postCollection.FindOneAndUpdate(s.ctx, bson.M{"_id": id}, bson.M{"$set": post})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

