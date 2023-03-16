package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var users []User

func init() {
	users = append(users, User{FirstName: "khushal", LastName: "Nagori"})
	users = append(users, User{FirstName: "test", LastName: "test"})
}

func (controller *UserController) GetUsers() {
	controller.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	controller.Data["json"] = users
	controller.ServeJSON()
}