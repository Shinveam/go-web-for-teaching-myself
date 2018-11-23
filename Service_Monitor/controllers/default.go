package controllers

import (
	"Service_Monitor/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"os/exec"
	"time"
)

type MainController struct {
	beego.Controller
}

//显示登录界面
func (c *MainController) ShowLogin() {
	//cookie的处理
	username := c.Ctx.GetCookie("username")//获得cookie值
	password := c.Ctx.GetCookie("password")
	if username != "" && password != "" {
		c.Data["username"] = username
		c.Data["password"] = password
		c.Data["check"] = "checked"
	}else {
		c.Data["check"] = "check"
	}
	c.TplName = "login.html"
}
//登录操作
func (c *MainController) Login() {
	username := c.GetString("username")
	pwd := c.GetString("password")

	//beego.Info(username, pwd)
	if username == "" || pwd == "" {
		beego.Info("用户没有输入")
		return
	}

	resp, err := models.SearchUser(username, pwd)
	if err != nil {
		beego.Info("查找失败：", err)
		return
	}
	if resp == nil {
		beego.Info("不存在该用户")
		c.Ctx.WriteString("failed")
		return
	}
	beego.Info("登录成功-->", username)
	//cookie的处理
	c.Ctx.SetCookie("username", username, 3600*time.Second)//设置cookie值，一小时后失效
	remember := c.GetString("remember")
	if remember == "on" {
		c.Ctx.SetCookie("password", pwd, 3600*time.Second)
	}else {
		c.Ctx.SetCookie("password", "123", -1)//设置负值可以删除cookie，此时第二个参数无效，可以随便填
	}

	//session的处理
	c.SetSession("username", username)//设置session
	c.SetSession("password", pwd)

	c.Ctx.WriteString("success")
}

//服务监控页面（主页）
func (c *MainController) ShowService() {
	//获取服务状态
	//err := models.ServiceStatus()
	//if err != nil {
	//	beego.Info("服务状态获取失败：", err)
	//	return
	//}

	//pageSize := 10//设置每页的条目数
	//pageIndexRaw := c.GetString("page")//获取页码
	//pageIndex, err := strconv.Atoi(pageIndexRaw)
	//if err != nil {
	//	pageIndex = 1 //首页
	//}

	Services, err := models.SearchService(0)
	if err != nil {
		beego.Info("服务查询出错：", err)
		return
	}
	Services=UseSSH(Services)//使用ssh查询服务状态
	//count := len(Services)
	//pageCount1 := float64(count)/float64(pageSize) //页码总数
	//pageCount := math.Ceil(pageCount1)//向上取整，如3/2=1.5，那么通过该函数，3/2=2；math.floor(···)向下取整
	////处理上下一页超出范围的问题
	//FirstPage := false//设置首页标记
	//if pageIndex == 1 {
	//	FirstPage = true //当pageIndex为首页时，标价为true
	//}
	//LastPage := false//设置末页标记
	//if pageIndex == int(pageCount) {
	//	LastPage = true//当pageIndex为末页时，标价为true
	//}
	//
	//c.Data["FirstPage"] = FirstPage
	//c.Data["LastPage"] = LastPage
	//c.Data["count"] = count
	//c.Data["pageCount"] = pageCount
	//c.Data["pageIndex"] = pageIndex

	//for i, v := range Services {
	//	serviceName := v.Name
	//	//cmd := exec.Command("netstat -lntup|grep " + serviceName +" |wc -l")
	//	fmt.Print(serviceName)
	//	num := 1//测试
	//	//num, _ := cmd.Output()
	//	fmt.Println(reflect.TypeOf(num))//测试
	//	status := strconv.Itoa(num)//测试
	//	//if  err != nil {
	//	//	return
	//	//}
	//
	//	//将状态保存在数据库
	//	if status == "0" {
	//		Services[i].Status = "down"
	//	}else {
	//		Services[i].Status = "up"
	//	}
	//	fmt.Println(Services[i].Status)
	//}

	c.Data["Services"] = Services
	c.TplName = "index.html"
}

//显示添加服务界面
func (c *MainController) ShowAddService() {
	c.TplName = "addservice.html"
}
//添加服务
func (c *MainController) AddService() {
	IpAddr := c.GetString("IpAddr")
	ServiceName := c.GetString("ServiceName")
	Email := c.GetString("Email")

	if IpAddr == "" || ServiceName == "" || Email == "" {
		beego.Info("输入内容为空！")
		return
	}

	resp, err := models.AddService(IpAddr, ServiceName, Email)
	if resp == 0 || err != nil {
		beego.Info("服务添加失败：", err)
		c.Ctx.WriteString("failed")
		return
	}
	c.Ctx.WriteString("success")
}
//删除服务
func (c *MainController) Delete()  {
	id, err := c.GetInt("id")
	fmt.Println(id)
	if err != nil {
		beego.Info("服务ID获取失败：", err)
		return
	}
	resp, err := models.DeleteService(id)
	if resp == 0 || err != nil {
		beego.Info("服务删除失败：", err)
		return
	}
	beego.Info("服务删除成功！")
}
//显示修改界面
func (c *MainController) ShowRevise()  {
	rid, err := c.GetInt("rid")
	if err != nil {
		beego.Info("服务ID获取失败：", err)
		return
	}
	services, err := models.SearchService(rid)//由于不需要对页面进行操作，所以pageSize和pageIndex可以为任意值
	if err != nil || services == nil {
		beego.Info("获取服务信息失败：", err)
		return
	}
	c.Data["services"] = services[0]
	c.TplName = "revise.html"
}
//修改服务
func (c *MainController) Revise()  {
	ServiceId, err := c.GetInt("ServiceId")
	if err != nil {
		beego.Info("修改服务--获取服务ID失败：", err)
		return
	}
	IpAddr := c.GetString("IpAddr")
	ServiceName := c.GetString("ServiceName")
	Email := c.GetString("Email")
	resp, err := models.ReviseService(ServiceId, IpAddr, ServiceName, Email)
	if err != nil || resp == 0 {
		beego.Info("修改失败：", err)
		c.Ctx.WriteString("failed")
		return
	}
	c.Ctx.WriteString("success")
}

func (c *MainController) Re() {
	Services, err := models.SearchService(0)
	if err != nil {
		beego.Info("服务查询出错：", err)
		return
	}
	Services = UseSSH(Services)//使用ssh查询服务状态


	jsone, _ := json.Marshal(&Services)
	//c.Data["Services"] = Services
	c.Ctx.WriteString(string(jsone))
}
