package routers

import (
	"article/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article/*", beego.BeforeRouter, FilterFunc)//路由过滤器，在执行路由之前进行判断

	beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")
	beego.Router("/register", &controllers.MainController{}, "get:ShowRegister;post:Register")
	beego.Router("/article/index", &controllers.MainController{}, "get:Index")
	beego.Router("/article/publish", &controllers.MainController{}, "get:ShowPublish;post:Publish")
	beego.Router("/article/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/article/delete", &controllers.MainController{}, "post:Delete")
	beego.Router("/article/upload", &controllers.MainController{}, "post:Upload")
	beego.Router("/article/exit", &controllers.MainController{}, "get:Exit")
	beego.Router("/article/chat", &controllers.MainController{}, "get:ShowChat;post:Chat")
	beego.Router("/article/revise", &controllers.MainController{}, "get:ShowRevise;post:Revise")
	beego.Router("/article/404", &controllers.MainController{}, "*:ShowFalse")//*代表无论使用何种方式，都访问ShowFalse函数
}
//结合路由过滤器，判断有session
var FilterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("username")
	if userName == nil {
		ctx.Redirect(302, "/")
	}
}