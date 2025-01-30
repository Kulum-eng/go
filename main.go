package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	myGin := gin.Default()


	/* productRepository := adapters.NewMySQLProductRepository()
	
	getAllProductsUseCase := application.NewGetAllProductsUseCase(productRepository)
	createProductUseCase := application.NewCreateProductUseCase(productRepository)

	getAllProductsController := controllers.NewGetAllProductsController(getAllProductsUseCase)
	createProductController := controllers.NewCreateProductController(createProductUseCase)

	routes.RegisterProductsRoutes(myGin, getAllProductsController, createProductController) */

	myGin.Run()
}
