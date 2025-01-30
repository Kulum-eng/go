package routes

import (
	"association/infrastructure/http/controllers"

	"github.com/gin-gonic/gin"

)

func SetupRoutes(router *gin.Engine, controller *controllers.Controller) {
	associationRoutes := router.Group("/associations")
	{
		associationRoutes.POST("/", controller.Create)
	}
}
