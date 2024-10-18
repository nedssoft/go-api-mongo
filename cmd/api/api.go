package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/routes"
)

type APIServer struct {
	db   *db.DB
	ctx  context.Context
}

func NewAPIServer(db *db.DB, ctx context.Context) *APIServer {
	return &APIServer{
		db: db,
		ctx: ctx,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	subrouter := router.Group("/api/v1")

	routes := routes.NewRoutes(subrouter, s.db, s.ctx)
	routes.RegisterRoutes()
	return router.Run()
}
