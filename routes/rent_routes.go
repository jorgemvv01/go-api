package routes

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/storage"
)

func RegisterRentRoutes(router *gin.RouterGroup) {
	db := storage.GetInstance()
	rentRepository := repositories.NewRentRepository(db)
	rentController := controllers.NewRentController(rentRepository)

	genreRouter := router.Group("/rent")
	genreRouter.POST("/create", rentController.Create)
}
