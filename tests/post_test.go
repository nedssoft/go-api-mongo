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
)

const (
	postTitle = "Test Post"
	postBody = "This is a test post"
)

func TestCreatePost(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)
	
	userCollection := db.GetCollection("users")
	user := models.User{Email: "test@example.com", Password: "password"}
	user.PreSave()
	userCollection.InsertOne(ctx, user)
	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)
	postPayload := requests.PostPayload{
		Title:  postTitle,
		Body:   postBody,
	}
	jsonValue, _ := json.Marshal(postPayload)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]responses.PostResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, postTitle, response["post"].Title)
	assert.Equal(t, postBody, response["post"].Body)
	assert.Equal(t, user.ID.Hex(), response["post"].UserID)
}

func TestGetPost(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)

	user := models.User{Email: "test@example.com", Password: "password"}
	user.PreSave()
	userCollection := db.GetCollection("users")
	userCollection.InsertOne(ctx, user)

	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)
	// Create a test post
	postCollection := db.GetCollection("posts")
	testPost := models.Post{Title: postTitle, Body: postBody, UserID: user.ID}
	testPost.PreSave()
	postCollection.InsertOne(ctx, testPost)

	req, _ := http.NewRequest("GET", "/api/v1/posts/"+testPost.ID.Hex(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]responses.PostResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, postTitle, response["post"].Title)
	assert.Equal(t, postBody, response["post"].Body)
	assert.Equal(t, user.ID.Hex(), response["post"].UserID)
}

func TestGetPosts(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)

	user := models.User{Email: "test@example.com", Password: "password"};user.PreSave()
	userCollection := db.GetCollection("users")
	userCollection.InsertOne(ctx, user)
	
	// Create test posts
	post1 := models.Post{Title: "Test Post 1", Body: "This is test post 1", UserID: user.ID}
	post1.PreSave()
	post2 := models.Post{Title: "Test Post 2", Body: "This is test post 2", UserID: user.ID}
	post2.PreSave()
	testPosts := []interface{}{
		post1,
		post2,
	}
	postCollection := db.GetCollection("posts")
	postCollection.InsertMany(ctx, testPosts)

	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)
	req, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]responses.PostResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response["posts"], 2)
	assert.Equal(t, "Test Post 1", response["posts"][0].Title)
	assert.Equal(t, "Test Post 2", response["posts"][1].Title)
}

func TestDeletePost(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)
	// Create a test post
	user := models.User{Email: "test@example.com", Password: "password"}
	user.PreSave()
	userCollection := db.GetCollection("users")
	
	userCollection.InsertOne(ctx, user)

	testPost := models.Post{Title: "Test Post", Body: "This is a test post", UserID: user.ID}
	testPost.PreSave()
	postCollection := db.GetCollection("posts")
	postCollection.InsertOne(ctx, testPost)

	req, _ := http.NewRequest("DELETE", "/api/v1/posts/"+testPost.ID.Hex(), nil)
	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Post deleted", response["message"])
}

func TestUpdatePost(t *testing.T) {
	router, db := SetupTestRouter()
	ctx := context.Background()
	defer CleanUp(db, ctx)
	// Create a test post
	user := models.User{Email: "test@example.com", Password: "password"}
	user.PreSave()
	userCollection := db.GetCollection("users")

	userCollection.InsertOne(ctx, user)
	testPost := models.Post{Title: postTitle, Body: postBody, UserID: user.ID}
	testPost.PreSave()
	postCollection := db.GetCollection("posts")
	postCollection.InsertOne(ctx, testPost)

	updatedPost := requests.PostPayload{Title: "Updated Post", Body: "This is an updated post"}
	jsonValue, _ := json.Marshal(updatedPost)
	req, _ := http.NewRequest("PUT", "/api/v1/posts/"+testPost.ID.Hex(), bytes.NewBufferString(string(jsonValue)))
	token, err := auth.NewJWTGenerator().GenerateToken(user.ID.Hex())
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]requests.PostPayload
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Updated Post", response["post"].Title)
	assert.Equal(t, "This is an updated post", response["post"].Body)
}
