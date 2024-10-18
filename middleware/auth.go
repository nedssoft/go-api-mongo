package middleware

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-api-mongo/auth"
	"github.com/nedssoft/go-api-mongo/service"
	"github.com/nedssoft/go-api-mongo/bin/db"
)

func AuthMiddleware(db *db.DB, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenGenerator := auth.NewJWTGenerator()
		userId, err := tokenGenerator.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		uid, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		user, err := service.NewUserService(db, ctx).GetUserById(uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()		
	}
}