package routers

import (
	"Beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{}, "get:GetUsers")
	beego.Router("/users", &controllers.UserController{}, "post:PostUser")
	beego.Router("/users/:username", &controllers.UserController{}, "get:GetUser")
}
