package routes

import (
	"association/infrastructure/http/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller *controllers.AssociationController) {
	association := router.Group("/associations")
	{
		association.POST("/", controller.Create)
		association.GET("/", controller.GetAll)
		association.GET("/:id", controller.GetByID)
		association.PUT("/:id", controller.Update)
		association.DELETE("/:id", controller.Delete)
	}
}
