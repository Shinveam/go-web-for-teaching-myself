package controllers

import (
	"article/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
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
	qs := o.QueryTable("Article")
	//qs.All(&articles)//select * from article

	count, err := qs.Count() //查询条目数
	pageSize := 10//设置每页的条目数

	/*
	pageIndex可以通过在前端使用/url?pageIndex=xxx的方式向后端传递
	后端的接受则使用pageIndex := c.GetString("pageIndex")获取
	将pageIndex转成整型，pageIndex, err := Strconv.Atoi(pageIndex)//首页的显示：if err != nil {pageIndex}
	上行的pageIndex在上上一行已经使用，这里只是作为提醒，与下行代码匹配
	*/
	pageIndex := 1//首页,由于前端的模板中暂时无法找到对应的首页（或上一页、下一页和末页）标签，故此处先使用1作为首页
	start := pageSize*(pageIndex-1)//起始位置
	qs.Limit(pageSize, start).All(&articles)//Limit()用于求每页的数据，第一个参数是每页显示多少条，第二个参数是起始位置

	pageCount1 := float64(count)/float64(pageSize) //页码总数
	pageCount := math.Ceil(pageCount1)//向上取整，如3/2=1.5，那么通过该函数，3/2=2；math.floor(···)向下取整

	if err != nil {
		beego.Info("查询出错：", err)
		return
	}
	c.Data["articles"] = articles

	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageIndex"] = pageIndex

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
	imgname := c.GetString("imgname")
	timer := time.Now()

	if artiname == "" || content == "" {
		beego.Info("非法输入！")
		return
	}
	resp, err := models.CreatArticle(artiname, content, timer, imgname)
	if err != nil || resp == 0 {
		beego.Info("文章创建失败！")
		c.Ctx.WriteString("failed")
		return
	}
	c.Ctx.WriteString("success")
}

//文件上传
func (c *MainController) Upload() {
	timer := time.Now()
	f, h, err := c.GetFile("uploadImg")

	//1、上传文件时如果文件同名则不再上传
	//2、如何处理？可以重命名文件名，使用原文件名加时间的组合构造新文件名
	//3、限定文件格式
	fileext := path.Ext(h.Filename) //获得文件格式
	if fileext != ".jpg" && fileext != ".png" {
		beego.Info("上传文件格式错误")
		return
	}
	//4、限制文件大小
	if h.Size > 5000000 {
		beego.Info("上传文件太大")
		return
	}
	//重命名文件名
	filename := timer.Format("2006-01-02-15-04-05") + fileext
	beego.Info(filename)
	defer f.Close()
	if err != nil {
		beego.Info("文件上传失败！")
		return
	}else {
		err := c.SaveToFile("uploadImg", "./static/userImg/"+filename)//路径前必须加“.”，否则无法识别
		if err != nil {
			beego.Info(err)
			c.Ctx.WriteString("failed")
		}
		c.Ctx.WriteString(filename)
	}
}

//文章内容详情
func (c *MainController) ShowContent() {
	c.TplName = "content.html"
}

//文章删除操作
func (c *MainController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章ID失败", err)
		return
	}
	err = models.DeleteArticle(id)
	if err != nil {
		beego.Info("文章删除失败", err)
		return
	}
	beego.Info("文章删除删除成功！")
}
