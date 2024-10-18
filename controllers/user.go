package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/data/requests"
	"github.com/nedssoft/go-api-mongo/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(db *db.DB, ctx context.Context) *UserController {
	return &UserController{
		UserService: service.NewUserService(db, ctx),
	}
}

func (c *UserController) CreateUser(gn *gin.Context) {
	var user *requests.UserPayload
	if err := gn.ShouldBindJSON(&user); err != nil {

		gn.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.HashPassword(); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser, err := c.UserService.CreateUser(user)
	if  err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gn.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func (c *UserController) GetUserWithPosts(gn *gin.Context) {
	id := gn.Param("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	result, err := c.UserService.GetUserWithPosts(uid)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"data": result})
}

func (c *UserController) GetUsers(gn *gin.Context) {
	users, err := c.UserService.GetUsers()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"users": users})
}

func (c *UserController) DeleteUser(gn *gin.Context) {
	id := gn.Param("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := c.UserService.DeleteUser(uid); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (c *UserController) UpdateUser(gn *gin.Context) {
	id := gn.Param("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user *requests.UserUpdatePayload
	if err := gn.ShouldBindJSON(&user); err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Failed to extract user payload"})
		return
	}
	if err := c.UserService.UpdateUser(uid, user); err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	gn.JSON(http.StatusOK, gin.H{"user": user})
}
