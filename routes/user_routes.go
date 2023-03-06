package routes

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/storage"
)

func RegisterUserRouter(router *gin.Engine) {
	db := storage.GetInstance()
	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	userRouter := router.Group("/users")
	userRouter.GET("", userController.GetAll)
	userRouter.GET("/:id", userController.GetByID)
	userRouter.POST("/create", userController.Create)
}
