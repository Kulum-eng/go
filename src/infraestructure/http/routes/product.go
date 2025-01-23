package routes

import (
	"github.com/gin-gonic/gin"

	"api/src/infraestructure/http/controllers"
)

func RegisterProductsRoutes(engine *gin.Engine, getAllProductsController *controllers.GetAllProductsController, createProductController *controllers.CreateProductController) {

	productsRoutes := engine.Group("/products")
	{
		productsRoutes.GET("", getAllProductsController.Run)
		productsRoutes.POST("", createProductController.Run)
	}
}
