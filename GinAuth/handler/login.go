package handler

import (
	"khushal/GinAuth/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user auth.LoginUser
	ctx.BindJSON(&user)
	token, err := auth.Auth(user)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, auth.Resposne{Message: token})
}

func Validate(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	resposne, err := auth.Validate(token)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resposne)
}
