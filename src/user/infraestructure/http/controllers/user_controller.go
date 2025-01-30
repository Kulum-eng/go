package controllers

import (
	"net/http"
	"user/application"
	"user/domain"
	"user/infrastructure/http/responses"

	"github.com/gin-gonic/gin"

)

type UserController struct {
	createUserUseCase *application.CreateUserUseCase
}

func NewUserController(createUC *application.CreateUserUseCase) *UserController {
	return &UserController{createUserUseCase: createUC}
}

func (ctrl *UserController) Create(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	if err := ctrl.createUserUseCase.Execute(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear usuario", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Usuario creado exitosamente", user))
}
