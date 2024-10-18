package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-api-mongo/bin/db"

	"github.com/nedssoft/go-api-mongo/controllers"
)

type AuthRoutes struct {
	db     *db.DB
	router gin.RouterGroup
	ctx context.Context
}

func NewAuthRoutes(router *gin.RouterGroup, db *db.DB, ctx context.Context) *AuthRoutes {
	return &AuthRoutes{
		db:     db,
		router: *router,
		ctx: ctx,
	}
}

func (r *AuthRoutes) RegisterRoutes() {
	authController := controllers.NewAuthController(r.db, r.ctx)
	r.router.POST("/login", authController.Login)
}
