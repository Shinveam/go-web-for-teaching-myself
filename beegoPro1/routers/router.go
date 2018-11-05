package routers

import (
	"beegoPro1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.MainController{}, "get:RegisterGet;post:RegisterPost")
    beego.Router("/login", &controllers.MainController{}, "get:LoginGet;post:LoginPost")
}
