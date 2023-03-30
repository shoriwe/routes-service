package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/models"
)

func (h *Handler) CreateVehicle(ctx *gin.Context) {
	var vehicle models.Vehicle
	bErr := ctx.Bind(&vehicle)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	cErr := h.Controller.CreateVehicle(&vehicle)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) DeleteVehicle(ctx *gin.Context) {
	vehicleUUID := ctx.Param(UUIDParam)
	cErr := h.Controller.DeleteVehicle(vehicleUUID)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) UpdateVehicle(ctx *gin.Context) {
	var vehicle models.Vehicle
	bErr := ctx.Bind(&vehicle)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	cErr := h.Controller.UpdateVehicle(&vehicle)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, StatusOK)
}

func (h *Handler) QueryVehicles(ctx *gin.Context) {
	var filter controller.VehicleFilter
	bErr := ctx.Bind(&filter)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	results, cErr := h.Controller.QueryVehicles(&filter)
	if cErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: cErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
