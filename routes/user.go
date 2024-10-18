package routes

import (
	"context"

	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/controllers"
	"github.com/nedssoft/go-api-mongo/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	db     *db.DB
	router gin.RouterGroup
	ctx    context.Context
}

func NewUserRoutes(router *gin.RouterGroup, db *db.DB, ctx context.Context) *UserRoutes {
	return &UserRoutes{
		db:     db,
		router: *router,
		ctx:    ctx,
	}
}

func (r *UserRoutes) RegisterRoutes() {
	userController := controllers.NewUserController(r.db, r.ctx)
	idRoute := r.router.Group("/users/:id")
	idRoute.GET("/posts", middleware.AuthMiddleware(r.db, r.ctx),userController.GetUserWithPosts)
	r.router.GET("/users", userController.GetUsers)
	r.router.POST("/users", userController.CreateUser)
	idRoute.DELETE("", userController.DeleteUser)
	idRoute.PATCH("", middleware.AuthMiddleware(r.db, r.ctx), userController.UpdateUser)
}
