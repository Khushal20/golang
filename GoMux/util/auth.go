package util

import (
	"khushal/goMux/entity"
	"khushal/goMux/service"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct{
	Email     string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret_key")

func GenerateToken(userLogin entity.UserLogin, service service.UserService) string{
	user := service.GetUserByEmail(userLogin.Email)
	if (user.Password != userLogin.Password){
		return "invalid user!"
	}
	exp_time := time.Now().Add(time.Minute*10) 

	claim:=Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp_time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) 
	token_str, err := token.SignedString(jwtKey) 
	if err!=nil{
		panic(err.Error())
	}
	return token_str
}

// func validate(token string) (bool, error){

// }