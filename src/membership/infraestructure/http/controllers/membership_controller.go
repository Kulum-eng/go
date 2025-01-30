package controllers

import (
	"membership/application"
	"membership/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

type MembershipController struct {
	create *application.CreateMembershipUseCase
	getU    *application.GetMembershipUseCase
	update *application.UpdateMembershipUseCase
	delete *application.DeleteMembershipUseCase
}

func NewMembershipController(
	create *application.CreateMembershipUseCase,
	get *application.GetMembershipUseCase,
	update *application.UpdateMembershipUseCase,
	delete *application.DeleteMembershipUseCase,
) *MembershipController {
	return &MembershipController{create, get, update, delete}
}

func (c *MembershipController) CreateMembership(ctx *gin.Context) {
	var membership domain.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if err := c.create.Execute(membership); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Membership created"})
}

func (c *MembershipController) GetMembership(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	membership, err := get.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, membership)
}

func (c *MembershipController) GetAllMemberships(ctx *gin.Context) {
	memberships, err := get.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, memberships)
}

func (c *MembershipController) UpdateMembership(ctx *gin.Context) {
	var membership domain.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if err := c.update.Execute(membership); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Membership updated"})
}

func (c *MembershipController) DeleteMembership(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.delete.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Membership deleted"})
}
