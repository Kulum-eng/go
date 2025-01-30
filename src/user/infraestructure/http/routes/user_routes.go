package routes

import (
	"user/infrastructure/http/controllers"

	"github.com/gin-gonic/gin"

)

func SetupUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controller.Create)
	}
}
