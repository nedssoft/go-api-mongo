package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nedssoft/go-api-mongo/auth"
	"github.com/nedssoft/go-api-mongo/data/requests"
	"github.com/nedssoft/go-api-mongo/data/responses"
	"github.com/nedssoft/go-api-mongo/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	name  = "Test User"
	email = "test@example.com"
	password = "password$123"
)

func TestCreateUser(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)

	user := requests.UserPayload{
		Name:  name,
		Email: email,
		Password: password,
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]responses.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, name, response["user"].Name)
	assert.Equal(t, email, response["user"].Email)
}

func TestGetUserWithPosts(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)

	user := models.User{
		Name:  name,
		Email: email,
		Password: password,
	}
	userCollection := db.GetCollection("users")
	user.PreSave()
	userCollection.InsertOne(ctx, user)
	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)

	postCollection := db.GetCollection("posts")
	testPost := models.Post{Title: postTitle, Body: postBody, UserID: user.ID}
	testPost.PreSave()
	postCollection.InsertOne(ctx, testPost)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/"+user.ID.Hex()+"/posts", nil)
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]responses.UserPostsResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	data := response["data"]
	assert.Equal(t, name, data.User.Name)
	assert.Equal(t, email, data.User.Email)
	assert.Equal(t, user.ID, data.User.ID)
	assert.Equal(t, postTitle, data.Posts[0].Title)
	assert.Equal(t, postBody, data.Posts[0].Body)
	assert.Equal(t, user.ID.Hex(), data.Posts[0].UserID)
}


func TestDeleteUser(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)
	// Create a test user
	testUser := models.User{Name: name, Email: email, Password: password}
	testUser.PreSave()
	userCollection := db.GetCollection("users")
	userCollection.InsertOne(ctx, testUser)
	req, _ := http.NewRequest("DELETE", "/api/v1/users/"+testUser.ID.Hex(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "User deleted", response["message"])

	// Check if the user was deleted
	var deletedUser models.User
	userCollection.FindOne(ctx, bson.M{"_id": testUser.ID}).Decode(&deletedUser)
	assert.Equal(t, primitive.NilObjectID, deletedUser.ID)
}

func TestUpdateUser(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)
	// Create a test user
	testUser := models.User{Name: name, Email: email, Password: password}
	testUser.PreSave()
	userCollection := db.GetCollection("users")
	userCollection.InsertOne(ctx, testUser)

	token, err := auth.NewJWTGenerator().GenerateToken(testUser.ID.Hex())
	assert.NoError(t, err)

	updatedUser := requests.UserUpdatePayload{Name: "Updated User"}
	jsonValue, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PATCH", "/api/v1/users/"+testUser.ID.Hex(), bytes.NewBufferString(string(jsonValue)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]responses.UserResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Updated User", response["user"].Name)
}

