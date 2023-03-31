package api

import (
	"encoding/base64"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/routes-service/controller"
	"github.com/shoriwe/routes-service/models"
)

type JWTResponse struct {
	JWT string `json:"jwt"`
}

var bearer = regexp.MustCompile(`(?m)^Bearer\s+`)

func (h *Handler) CheckAPIKey(ctx *gin.Context) {
	tokenBase64 := bearer.ReplaceAllString(ctx.GetHeader("Authorization"), "")
	tokenBytes, dErr := base64.StdEncoding.DecodeString(tokenBase64)
	if dErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Status{Succeed: false, Error: dErr.Error()})
		return
	}
	apiKey, aErr := h.Controller.AuthorizeAPIKey(string(tokenBytes))
	if aErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Status{Succeed: false, Error: aErr.Error()})
		return
	}
	ctx.Set(APIKeyKey, apiKey)
	ctx.Next()
}

func (h *Handler) CheckJWT(ctx *gin.Context) {
	tokenBase64 := bearer.ReplaceAllString(ctx.GetHeader("Authorization"), "")
	tokenBytes, dErr := base64.StdEncoding.DecodeString(tokenBase64)
	if dErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Status{Succeed: false, Error: dErr.Error()})
		return
	}
	credentials, aErr := h.Controller.AuthorizeUser(string(tokenBytes))
	if aErr != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Status{Succeed: false, Error: aErr.Error()})
		return
	}
	ctx.Set(CredentialsKey, credentials)
	ctx.Next()
}

func (h *Handler) OnlyAdmin(ctx *gin.Context) {
	user := ctx.MustGet(CredentialsKey).(*models.User)
	if user.Type != models.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, ForbiddenStatus)
		return
	}
	ctx.Next()
}

func (h *Handler) Login(ctx *gin.Context) {
	var credentials models.User
	bErr := ctx.Bind(&credentials)
	if bErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Status{Succeed: false, Error: bErr.Error()})
		return
	}
	token, loginErr := h.Controller.Login(&credentials)
	switch loginErr {
	case nil:
		ctx.JSON(http.StatusOK, JWTResponse{JWT: token})
	case controller.ErrorUnauthorized:
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedStatus)
	default:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, loginErr.Error())
	}
}
