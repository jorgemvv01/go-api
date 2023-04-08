package storage

import (
	"github/jorgemvv01/go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
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

func MigrateModels(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Rent{},
		&models.Type{},
		&models.Genre{},
		&models.Movie{},
		&models.MovieRent{},
	); err != nil {
		panic("failed to migrate models")
	}

	tx := db.Begin()

	var movieTypes = []models.Type{
		{Name: "New releases"},
		{Name: "Regular movies"},
		{Name: "Old movies"},
	}

	for _, t := range movieTypes {
		if err := tx.Create(&t).Error; err != nil {
			tx.Rollback()
			panic("failed to create movie types")
		}
	}

	if err := tx.Commit().Error; err != nil {
		panic("failed to commit types transaction")
	}

}
