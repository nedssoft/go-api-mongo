package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-api-mongo/bin/db"
)

type Routes struct {
	db     *db.DB
	router gin.RouterGroup
	ctx    context.Context
}

func NewRoutes(router *gin.RouterGroup, db *db.DB, ctx context.Context) *Routes {
	return &Routes{
		db:     db,
		router: *router,
		ctx:    ctx,
	}
}

func (r *Routes) RegisterRoutes() {
	postRoutes := NewPostRoutes(&r.router, r.db, r.ctx)
	postRoutes.RegisterRoutes()

	userRoutes := NewUserRoutes(&r.router, r.db, r.ctx)
	userRoutes.RegisterRoutes()

	authRoutes := NewAuthRoutes(&r.router, r.db, r.ctx)
	authRoutes.RegisterRoutes()
}
