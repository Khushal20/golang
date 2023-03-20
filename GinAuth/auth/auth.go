package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var key = []byte("TestKey")
var logger = log.Default()

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Resposne struct {
	Message string `json:"message"`
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

var Users = map[string]string{
	"khushal": "password",
	"knagori": "Test123.",
}

func Auth(user LoginUser) (string, error) {
	password, ok := Users[user.UserName]
	if !ok {
		logger.Fatal("Invalid user", user.UserName)
		return "", fmt.Errorf("Invalid user %v", user.UserName)
	}
	if password != user.Password {
		logger.Fatal("Invalid password", user.UserName)
	}
	exp := time.Now().Add(time.Minute * 10)
	claim := &Claims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(key)
	if err != nil {
		logger.Println("Error while getting token", user.UserName)
		return "", err
	}
	return tokenString, nil
}

func Validate(tokenString string) (Resposne, error) {
	var claim Claims
	_, err := validator(tokenString, &claim)

	if err != nil {
		logger.Println("Invalid token")
		return Resposne{Message: "Invalid token"}, err
	}

	err = claim.Valid()
	if err != nil {
		logger.Println("Invalid token")
		return Resposne{Message: "Invalid token"}, err
	}
	return Resposne{Message: fmt.Sprintf("Welcome %v", claim.Username)}, nil
}

func validator(tokenString string, claim *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		return key, nil
	})
}


func Add(user LoginUser){
	Users[user.UserName] = user.Password;
}