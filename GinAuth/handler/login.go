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

func ValidateAPi(ctx *gin.Context) {
	Validate(ctx)
	ctx.AbortWithStatus(http.StatusOK)
}

func Validate(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	token := strings.ReplaceAll(bearer, "Bearer ", "")
	_, err := auth.Validate(token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}
	ctx.Next()
}

type User struct{
	UserName string `json:"name"`
	Password string `json:"password"`
}


func Userget(ctx *gin.Context) {
	var users []User
	for un, pass := range auth.Users{
		users = append(users, User{UserName: un, Password: pass})
	}
	ctx.AbortWithStatusJSON(http.StatusOK, users)
}