package storage

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	if db == nil {
		var err error
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}
