package controllers

import (
	"article/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type MainController struct {
	beego.Controller
}

//显示登录界面
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}
//登录操作
func (c *MainController) Login() {
	username := c.GetString("username")
	pwd := c.GetString("password")
	beego.Info(username, pwd)

	if username == "" || pwd == "" {
		beego.Info("用户没有输入")
		return
	}

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
func (c *MainController) ShowRegister() {
	c.TplName = "register.html"
}
//注册操作
func (c *MainController) Register() {
	username := c.GetString("username")
	pwd := c.GetString("password")

	if username == "" || pwd == "" {
		beego.Info("用户没有输入")
		return
	}

	num, err := models.RegisterUser(username, pwd)
	if err != nil {
		beego.Info("注册失败")
		c.Ctx.WriteString("failed")
		return
	}
	if num == 0 {
		beego.Info("用户已存在")
		c.Ctx.WriteString("failed")
		return
	}
	beego.Info("注册成功")
	c.Ctx.WriteString("success")
}

//显示主页
func (c *MainController) Index() {
	//显示文章标题、浏览量、创建时间到index.html中
	//1、查询
	o := orm.NewOrm()
	var articles []models.Article
	_, err := o.QueryTable("Article").All(&articles)
	if err != nil {
		beego.Info("查询出错：", err)
		return
	}
	beego.Info(articles)
	c.Data["articles"] = articles

	c.TplName = "index.html"
}

//显示文章创建页
func (c *MainController) ShowPublish() {
	c.TplName = "publish.html"
}

//创建文章操作
func (c *MainController) Publish() {
	artiname := c.GetString("artiname")
	content := c.GetString("content")
	if artiname == "" || content == "" {
		beego.Info("非法输入！")
		return
	}
	timer := time.Now()
	resp, err := models.CreatArticle(artiname, content, timer)
	if err != nil || resp == 0 {
		beego.Info("文章创建失败！")
		c.Ctx.WriteString("failed")
		return
	}
	c.Ctx.WriteString("success")
}

//文章内容详情
func (c *MainController) ShowContent() {
	c.TplName = "content.html"
}
