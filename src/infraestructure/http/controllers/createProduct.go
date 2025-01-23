package controllers

import (
	"api/src/application"
	"api/src/domain"
	"api/src/infraestructure/http/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	createProductUseCase *application.CreateProductUseCase
}

func NewCreateProductController(createProductUseCase *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		createProductUseCase: createProductUseCase,
	}
}

func (ctrl CreateProductController) Run(c *gin.Context) {
	var newProduct domain.Product

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.GeneralResponse{
			Success: false,
			Message: "ocurrió un error al leer la petición",
			Error:   err.Error(),
		})
		return
	}

	err = ctrl.createProductUseCase.Run(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.GeneralResponse{
			Success: false,
			Message: "ocurrió un error al guardar el producto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.GeneralResponse{
		Success: true,
		Message: "producto creado exitosamente",
		Data: newProduct,
	})
}