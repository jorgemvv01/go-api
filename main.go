package main

import (
	_ "github/jorgemvv01/go-api/docs"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/routes"
	"github/jorgemvv01/go-api/storage"
	"log"
)

// @title VideoClub GO-API
// @version 1.0
// @description A simple API to manage movie rentals

// @BasePath /api
func main() {

	db := storage.GetInstance()
	db.AutoMigrate(
		&models.User{},
		&models.Rent{},
		&models.Type{},
		&models.Genre{},
		&models.Movie{},
		&models.RentMovie{},
	)

	r := routes.SetupRoutes()
	log.Println("[--->>>> STARTING SERVER... <<<<---]")
	if err := r.Run(); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
