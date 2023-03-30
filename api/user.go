package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/models"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var user models.User
	bErr := ctx.Bind(&user)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	cErr := h.Controller.CreateUser(&user)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userUUID := ctx.Param(UUIDParam)
	cErr := h.Controller.DeleteUser(userUUID)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var user models.User
	bErr := ctx.Bind(&user)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	cErr := h.Controller.UpdateUser(&user)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) QueryUsers(ctx *gin.Context) {
	var filter controller.UserFilter
	bErr := ctx.Bind(&filter)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	results, cErr := h.Controller.QueryUsers(&filter)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
