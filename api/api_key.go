package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/models"
)

func (h *Handler) CreateAPIKey(ctx *gin.Context) {
	var ak models.APIKey
	bErr := ctx.Bind(&ak)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	cErr := h.Controller.CreateAPIKey(&ak)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) DeleteAPIKey(ctx *gin.Context) {
	akUUID := ctx.Param(UUIDParam)
	cErr := h.Controller.DeleteAPIKey(akUUID)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) QueryAPIKeys(ctx *gin.Context) {
	var filter controller.APIKeyFilter
	bErr := ctx.Bind(&filter)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	results, cErr := h.Controller.QueryAPIKeys(&filter)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
