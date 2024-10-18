package main

import (
	"context"
	"log"

	"github.com/nedssoft/go-api-mongo/bin/db"
	"github.com/nedssoft/go-api-mongo/cmd/api"
	"github.com/nedssoft/go-api-mongo/config"
)

func main() {
	ctx := context.Background()
	dbConfig := config.GetDBConfig()
	db, err := db.NewDB(dbConfig.MongoURI, dbConfig.DBName)
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(db, ctx)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
