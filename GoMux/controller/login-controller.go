package controller

import (
	"encoding/json"
	"khushal/goMux/entity"
	"khushal/goMux/service"
	"khushal/goMux/util"
	"net/http"
)

type LoginController interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type loginController struct{
	service service.UserService
}

type Token struct{
	token string
}

func NewLoginController(service service.UserService) *loginController{
	return &loginController{service}
}

func (controller loginController) Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var userLogin entity.UserLogin
	json.NewDecoder(r.Body).Decode(&userLogin)
	token := util.GenerateToken(userLogin, controller.service)
	json.NewEncoder(w).Encode(&Token{token: token})
}