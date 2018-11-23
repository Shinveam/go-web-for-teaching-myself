package routers

import (
	"Service_Monitor/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/service/*", beego.BeforeRouter, FilterFunc)//路由过滤器，在执行路由之前进行判断

	beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")
    beego.Router("/service/index", &controllers.MainController{}, "get:ShowService")
    beego.Router("/service/addservice", &controllers.MainController{}, "get:ShowAddService;post:AddService")
    beego.Router("/service/delete", &controllers.MainController{}, "get:Delete")
	beego.Router("/service/revise", &controllers.MainController{}, "get:ShowRevise;post:Revise")
	beego.Router("/service/re", &controllers.MainController{}, "post:Re")
}

//结合路由过滤器，判断有session
var FilterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("username")
	fmt.Println(userName)
	if userName == nil {
		ctx.Redirect(302, "/")
	}
}