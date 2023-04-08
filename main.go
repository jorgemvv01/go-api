package main

import (
	_ "github/jorgemvv01/go-api/docs"
	"github/jorgemvv01/go-api/routes"
	"github/jorgemvv01/go-api/storage"
	"log"
)

// @title VideoClub / Go-REST-API
// @version 1.0
// @description A simple Go-REST-API to manage movie rentals \n GitHub Repository: https://github.com/jorgemvv01/go-api

// @contact.name   Jorge Mario Villarreal V.
// @contact.url    https://www.linkedin.com/in/jorgemariovillarreal/
// @contact.email  jorgemvv01@gmail.com

// @BasePath /api
func main() {

	db := storage.GetInstance()
	storage.MigrateModels(db)

	r := routes.SetupRoutes()
	log.Println("[--->>>> STARTING SERVER... <<<<---]")
	if err := r.Run(); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
