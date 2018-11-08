package routers

import (
	"article/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterFunc)//路由过滤器，在执行路由之前进行判断

	beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")
	beego.Router("/register", &controllers.MainController{}, "get:ShowRegister;post:Register")
	beego.Router("/index", &controllers.MainController{}, "get:Index")
	beego.Router("/publish", &controllers.MainController{}, "get:ShowPublish;post:Publish")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/delete", &controllers.MainController{}, "post:Delete")
	beego.Router("/upload", &controllers.MainController{}, "post:Upload")
	beego.Router("/exit", &controllers.MainController{}, "get:Exit")
}
//结合路由过滤器，判断有session
//var FilterFunc = func(ctx *context.Context) {
//	userName := ctx.Input.Session("username")
//	if userName == nil {
//		ctx.Redirect(302, "/")
//	}
//}