package routers

import (
	"article/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")
	beego.Router("/register", &controllers.MainController{}, "get:ShowRegister;post:Register")
	beego.Router("/index", &controllers.MainController{}, "get:Index")
	beego.Router("/publish", &controllers.MainController{}, "get:ShowPublish;post:Publish")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/delete", &controllers.MainController{}, "post:Delete")
	beego.Router("/upload", &controllers.MainController{}, "post:Upload")
}
