package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Map(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Controller.Map)
}
