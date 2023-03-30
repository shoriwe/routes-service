package api

import "github.com/gin-gonic/gin"

func (h *Handler) Echo(ctx *gin.Context) {
	ctx.Done()
}
