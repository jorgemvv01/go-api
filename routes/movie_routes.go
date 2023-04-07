package routes

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/storage"
)

func RegisterMovieRouter(router *gin.RouterGroup) {
	db := storage.GetInstance()
	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	movieRouter := router.Group("/movies")
	movieRouter.GET("", movieController.GetAll)
	movieRouter.GET("/:id", movieController.GetByID)
	movieRouter.POST("/create", movieController.Create)
	movieRouter.PUT("/update/:id", movieController.Update)
	movieRouter.DELETE("/delete/:id", movieController.Delete)
}
