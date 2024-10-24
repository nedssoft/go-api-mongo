package tests

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nedssoft/go-api-mongo/routes"
	database "github.com/nedssoft/go-api-mongo/bin/db"
)

func SetupTestRouter() (*gin.Engine, *database.DB) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	ctx := context.Background()
  
	api := router.Group("/api/v1")
	db, err := database.NewDB("mongodb://localhost:27017", "test-go", ctx)
	if err != nil {
		log.Fatal(err)
	}

	routes := routes.NewRoutes(api, db, ctx)
	routes.RegisterRoutes()
	return router, db
}

func CleanUp(db *database.DB, ctx context.Context) {
	db.Db.Drop(ctx)
}
