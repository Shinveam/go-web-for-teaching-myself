package routers

import (
	"article/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article/*", beego.BeforeRouter, FilterFunc)//路由过滤器，在执行路由之前进行判断

	beego.Router("/", &controllers.MainController{}, "get:ShowLogin;post:Login")//登录
	beego.Router("/register", &controllers.MainController{}, "get:ShowRegister;post:Register")//注册
	beego.Router("/article/index", &controllers.MainController{}, "get:Index")//主页
	beego.Router("/article/publish", &controllers.MainController{}, "get:ShowPublish;post:Publish")//创建文章
	beego.Router("/article/content", &controllers.MainController{}, "get:ShowContent")//文章详情
	beego.Router("/article/delete", &controllers.MainController{}, "post:Delete")//文章删除
	beego.Router("/article/upload", &controllers.MainController{}, "post:Upload")//文件上传
	beego.Router("/article/exit", &controllers.MainController{}, "get:Exit")//退出
	beego.Router("/article/chat", &controllers.MainController{}, "get:ShowChat;post:Chat")//聊天
	beego.Router("/article/revise", &controllers.MainController{}, "get:ShowRevise;post:Revise")//文章修改
	beego.Router("/article/classify", &controllers.MainController{}, "get:ShowClassify")//分类文章显示
	beego.Router("/article/gallery", &controllers.MainController{}, "get:Gallery")//画廊
	beego.Router("/article/download", &controllers.MainController{}, "get:Download")//文件下载
	beego.Router("/article/404", &controllers.MainController{}, "*:ShowFalse")//*代表无论使用何种方式，都访问ShowFalse函数
}

//结合路由过滤器，判断有session
var FilterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("username")
	if userName == nil {
		ctx.Redirect(302, "/")
	}
}