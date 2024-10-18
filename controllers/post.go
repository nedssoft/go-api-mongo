package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/data/requests"
	"github.com/nedssoft/go-api-mongo/models"
	"github.com/nedssoft/go-api-mongo/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostController struct {
	PostService *service.PostService
}

func NewPostController(db *db.DB, ctx context.Context) *PostController {
	return &PostController{
		PostService: service.NewPostService(db, ctx),
	}
}

func (c *PostController) CreatePost(gn *gin.Context) {
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Error getting post's user"})
		return
	}
	var post requests.PostPayload
	if err := gn.BindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postResponse, err := c.PostService.CreatePost(&post, user.ID)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
	}
	gn.JSON(http.StatusCreated, gin.H{"post": postResponse})
}

func (c *PostController) GetPost(gn *gin.Context) {
	id := gn.Param("id")
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	post, err := c.PostService.GetPost(pid)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"post": post})
}

func (c *PostController) GetPosts(gn *gin.Context) {
	posts, err := c.PostService.GetPosts()
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get posts"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (c *PostController) DeletePost(gn *gin.Context) {
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		log.Println("user not found")
		gn.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	id := gn.Param("id")
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	post, err := c.PostService.GetPost(pid)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}
	if post.UserID != user.ID.Hex() {
		gn.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this post"})
		return
	}
	if err := c.PostService.DeletePost(pid); err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func (c *PostController) UpdatePost(gn *gin.Context) {
	id := gn.Param("id")
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	var post *requests.PostUpdatePayload
	if err := gn.BindJSON(&post); err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, ok := gn.MustGet("user").(*models.User)
	if !ok {
		log.Println("user not found")
		gn.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	if err := c.PostService.UpdatePost(pid, post, user.ID); err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"post": post})
}
