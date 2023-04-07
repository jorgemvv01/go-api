package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		RegisterUserRouter(api)
		RegisterTypeRoutes(api)
		RegisterGenreRoutes(api)
		RegisterMovieRouter(api)
	}

	return router
}
