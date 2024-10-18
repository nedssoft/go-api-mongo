package config

import (
	"os"
	 "github.com/lpernett/godotenv"
)

type DBConfig struct {
	MongoURI string
	DBName string
}

func GetDBConfig() *DBConfig {
	godotenv.Load()
  return &DBConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		DBName: os.Getenv("DB_NAME"),
	}
}
