package controllers

import (
	"api/src/association/application"
	"api/src/association/domain"
	"api/src/association/infraestructure/http/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssociationController struct {
	createAssociationUseCase *application.CreateAssociationUseCase
}

func NewAssociationController(createUC *application.CreateAssociationUseCase) *AssociationController {
	return &AssociationController{createAssociationUseCase: createUC}
}

func (ctrl *AssociationController) Create(ctx *gin.Context) {
	var association domain.Association
	if err := ctx.ShouldBindJSON(&association); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	if err := ctrl.createAssociationUseCase.Execute(association); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Asociación creada exitosamente", association))
}
