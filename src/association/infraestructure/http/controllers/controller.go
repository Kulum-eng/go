package controllers

import (
	"association/application"
	"association/domain"
	"association/infrastructure/http/responses"
	"net/http"

	"github.com/gin-gonic/gin"

)

type Controller struct {
	useCase *application.CreateUseCase
}

func NewController(useCase *application.CreateUseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (ctrl *Controller) Create(ctx *gin.Context) {
	var association domain.Association
	if err := ctx.ShouldBindJSON(&association); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	if err := ctrl.useCase.Execute(association); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear la asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Asociación creada exitosamente", association))
}
