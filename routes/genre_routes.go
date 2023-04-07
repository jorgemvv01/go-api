package routes

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/storage"
)

func RegisterGenreRoutes(router *gin.RouterGroup) {
	db := storage.GetInstance()
	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	genreRouter := router.Group("/genres")
	genreRouter.GET("/", genreController.GetAll)
	genreRouter.GET("/:id", genreController.GetByID)
	genreRouter.POST("/create", genreController.Create)
	genreRouter.PUT("/update/:id", genreController.Update)
	genreRouter.DELETE("/delete/:id", genreController.Delete)
}
