package controllers

import (
	"beegoPro1/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//ORM数据插入操作
	////1、有ORM对象
	//o := orm.NewOrm()
	////2、有插入数据对象的结构体
	//user := models.User{}
	////3、对结构体赋值
	//user.Name = "wei"
	//user.Pwd = "123456"
	////4、插入
	//_, err := o.Insert(&user)
	//if err != nil {
	//	beego.Info("插入失败：", err)
	//	return  //插入失败要返回，不能继续插入往下走
	//}

	//ORM数据（简单）查询操作
	//1、有ORM对象
	o := orm.NewOrm()
	//2、有查询对象
	user := models.User{}
	//3、指定查询对象字段值
	user.Id = 1 //根据name查：user.Name = "wei"
	//4、查询
	err := o.Read(&user) //根据name查：err := o.Read(&user, "name")
	if err != nil {
		beego.Info("查询失败：", err)
		return
	}
	//显示查询成功信息
	beego.Info("查询成功：", user)

	//ORM更新操作
	////1、有ORM对象
	//o := orm.NewOrm()
	////2、有更新对象
	//user := models.User{}
	////3、查找需更新数据
	//user.Id = 1
	//err := o.Read(&user)
	////4、给数据重新赋值
	//if err ==nil {
	//	user.Pwd = "888888"
	//
	//	//5、更新
	//	_, err := o.Update(&user)
	//	if err != nil {
	//		beego.Info("更新失败：", err)
	//		return
	//	}
	//	beego.Info("更新成功：", user)
	//}

	//ORM删除操作
	////1、有ORM对象
	//o := orm.NewOrm()
	////2、有删除对象
	//user := models.User{}
	////3、指定删除哪一条数据
	//user.Id = 1
	////4、删除
	//_, err := o.Delete(&user)
	//if err != nil {
	//	beego.Info("删除错误：", err)
	//	return
	//}


	c.Data["Website"] = "*****"
	c.Data["Email"] = "xinwei1210@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController)RegisterGet()  {
	c.TplName = "signup1.html"
}

func (c *MainController)RegisterPost()  {
	//1、注册业务的实现
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info(userName, pwd)
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空！！！")
		c.Redirect("/register", 302) //重定向函数，重新返回register模板
		return
	}
	//···此处省略对数据库的操作代码
	c.Ctx.WriteString("注册成功")
}

func (c *MainController)LoginGet()  {
	c.TplName = "login1.html"
}

func (c *MainController)LoginPost()  {
	//登录业务的实现
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info(userName, pwd)
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空！！！")
		c.Redirect("/login", 302) //重定向函数，重新返回register模板
		return
	}
	//···此处省略登录业务的实现
	c.Ctx.WriteString("登录成功")
}


