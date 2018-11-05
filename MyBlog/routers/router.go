package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")
    beego.Router("/register", &controllers.MainController{}, "get:ShowRegister;post:Register")
    beego.Router("/index", &controllers.MainController{}, "get:Index")
}
