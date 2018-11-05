package controllers

import (
	"MyBlog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//显示登录界面
func (c *MainController) ShowLogin() {
	c.TplName = "login1.html"
}
//登录操作
func (c *MainController) Login()  {
	username := c.GetString("username")
	pwd := c.GetString("password")
	beego.Info(username, pwd)

	resp, err := models.SearchUser(username, pwd)
	if err != nil {
		beego.Info("查找失败")
		return
	}
	if resp == nil {
		beego.Info("不存在该用户")
		c.Ctx.WriteString("failed")
		return
	}
	beego.Info("登录成功")
	c.Ctx.WriteString("success")
	//c.Ctx.Redirect(302, "/index")
}

//显示注册界面
func (c *MainController) ShowRegister()  {
	c.TplName = "signup1.html"
}
//注册操作
func (c *MainController) Register() {
	username := c.GetString("username")
	pwd := c.GetString("password")

	resp, err := models.RegisterUser(username, pwd)
	if err != nil {
		beego.Info("注册失败")
	}
	if resp != nil {
		beego.Info("用户已存在")
	}
	beego.Info("注册成功")
	c.Ctx.Redirect(302, "/login")
}

//显示主页
func (c *MainController) Index()  {
	c.TplName = "index.html"
}
