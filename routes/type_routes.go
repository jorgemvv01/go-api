package routes

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/storage"
)

func RegisterTypeRoutes(router *gin.RouterGroup) {
	db := storage.GetInstance()
	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	typeRouter := router.Group("/types")
	typeRouter.GET("/", typeController.GetAll)
	typeRouter.GET("/:id", typeController.GetByID)
	//typeRouter.POST("/create", typeController.Create)
	//typeRouter.PUT("/update/:id", typeController.Update)
	//typeRouter.DELETE("/delete/:id", typeController.Delete)
}
