package routes

import (
	"context"
	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/controllers"
	"github.com/nedssoft/go-api-mongo/middleware"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	db     *db.DB
	router gin.RouterGroup
	ctx    context.Context
}

func NewPostRoutes(router *gin.RouterGroup, db *db.DB, ctx context.Context) *PostRoutes {
	return &PostRoutes{
		db:     db,
		router: *router,
		ctx:    ctx,
	}
}

func (r *PostRoutes) RegisterRoutes() {
	postController := controllers.NewPostController(r.db, r.ctx)
	idRoute := r.router.Group("/posts/:id")
	r.router.GET("/posts", middleware.AuthMiddleware(r.db, r.ctx), postController.GetPosts)
	idRoute.GET("", middleware.AuthMiddleware(r.db, r.ctx), postController.GetPost)
	r.router.POST("/posts", middleware.AuthMiddleware(r.db, r.ctx), postController.CreatePost)
	idRoute.DELETE("", middleware.AuthMiddleware(r.db, r.ctx), postController.DeletePost)
	idRoute.PUT("", middleware.AuthMiddleware(r.db, r.ctx), postController.UpdatePost)
}