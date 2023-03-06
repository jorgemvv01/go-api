package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	RegisterUserRouter(r)
	return r
}
