package controllers

import (
	"article/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
	"turingAPI/turing"
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
	//beego.Info("remember-->", remember)
	if remember == "on" {
		c.Ctx.SetCookie("password", pwd, 3600*time.Second)
	}else {
		c.Ctx.SetCookie("password", "123", -1)//设置负值可以删除cookie，此时第二个参数无效，可以随便填
	}

	//session的处理
	c.SetSession("username", username)//设置session

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
		beego.Info("注册失败：", err)
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
	//session处理
	//userName := c.GetSession("username")
	//if userName == nil {
	//	c.Redirect("/", 302)
	//	return
	//}
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article").RelatedSel("ArticleType")//关联ArticleType数据库，可以在多处添加RelatedSel
	//qs.All(&articles)//select * from article

	count, err := qs.Count() //查询条目数
	if err != nil {
		beego.Info("查询出错：", err)
		return
	}
	pageSize := 10//设置每页的条目数

	//首页、上一页、下一页和末页的处理
	/*
	pageIndex可以通过在前端使用/url?pageIndex=xxx的方式向后端传递
	后端的接受则使用pageIndex := c.GetString("pageIndex")获取
	将pageIndex转成整型，pageIndex, err := Strconv.Atoi(pageIndex)//首页的显示：if err != nil {pageIndex = 1}
	上行的pageIndex在上上一行已经使用，这里只是作为提醒，与下行代码匹配
	*/
	pageIndexRaw := c.GetString("page")//获取页码
	pageIndex, err := strconv.Atoi(pageIndexRaw)
	if err != nil {
		pageIndex = 1 //首页
	}
	start := pageSize*(pageIndex-1)//起始位置
	qs.Limit(pageSize, start).All(&articles)//Limit()用于求每页的数据，第一个参数是每页显示多少条，第二个参数是起始位置

	pageCount1 := float64(count)/float64(pageSize) //页码总数
	pageCount := math.Ceil(pageCount1)//向上取整，如3/2=1.5，那么通过该函数，3/2=2；math.floor(···)向下取整
	//处理上下一页超出范围的问题
	FirstPage := false//设置首页标记
	if pageIndex == 1 {
		FirstPage = true //当pageIndex为首页时，标价为true
	}
	LastPage := false//设置末页标记
	if pageIndex == int(pageCount) {
		LastPage = true//当pageIndex为末页时，标价为true
	}

	c.Data["FirstPage"] = FirstPage
	c.Data["LastPage"] = LastPage
	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["pageIndex"] = pageIndex

	c.TplName = "index.html"
}

//显示文章创建页
func (c *MainController) ShowPublish() {
	artiType, err := models.GetArtiType()
	if err != nil {
		beego.Info("类型查找失败：", err)
		return
	}
	c.Data["ArtiType"] = artiType
	c.TplName = "publish.html"
}
//创建文章操作
func (c *MainController) Publish() {
	artiname := c.GetString("artiname")
	content := c.GetString("content")
	imgname := c.GetString("imgname")
	timer := time.Now()

	typeId, err := c.GetInt("typeId")
	//beego.Info("TypeId:", typeId)
	if err != nil {
		beego.Info("下拉类型获取失败：", err)
		return
	}
	var artiType models.ArticleType
	artiType.Id = typeId

	if artiname == "" || content == "" {
		beego.Info("非法输入！")
		return
	}
	resp, err := models.CreatArticle(artiname, content, timer, imgname, artiType)
	if err != nil || resp == 0 {
		beego.Info("文章创建失败：", err)
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
		beego.Info("文件上传失败：", err)
		return
	}else {
		err := c.SaveToFile("uploadImg", "./static/userImg/"+filename)//路径前必须加“.”，否则无法识别
		if err != nil {
			beego.Info("文件保存失败：", err)
			c.Ctx.WriteString("failed")
		}
		c.Ctx.WriteString(filename)
	}
}

//文章内容详情
func (c *MainController) ShowContent() {
	o := orm.NewOrm()
	var articles []models.Article
	//侧边栏分类显示
	//1、类别名称显示
	artiType, err := models.GetArtiType()

	if err != nil {
		beego.Info("类型查找失败：", err)
		return
	}

	//文章详情显示
	artiId, err := c.GetInt("cid")
	if err != nil {
		beego.Info("获取文章ID失败：", err)
		return
	}
	_, err = o.QueryTable("Article").Filter("Id", artiId).All(&articles)
	if err != nil {
		beego.Info("文章内容获取失败：", err)
		return
	}
	//*获取文章图片


	//浏览量计数
	_, err = o.Raw("update article set count=count+1 where Id=?", artiId).Values(&[]orm.Params{})
	if err != nil {
		beego.Info("浏览量获取失败：", err)
		return
	}

	//获取文章类型下的文章数量
	for i, v:= range artiType {
		//1、获取文章类型的id
		//beego.Info("v.Id-->", v.Id)
		//beego.Info("v.Id.type-->", reflect.TypeOf(v.Id))

		//2、根据文章类型id检索article表的该id的数量
		//方式1：原生字符串查询
		//var maps []orm.Params
		//_, err := o.Raw("select count(*) from article where article_type_id=?", v.Id).Values(&maps)
		//beego.Info("count(*)-->",maps[0]["count(*)"])
		////方式二：ORM查询
		num, err := o.QueryTable("Article").Filter("ArticleType__id", v.Id).RelatedSel().Count()
		//beego.Info("num-->", num)
		if err != nil {
			beego.Info("读取文章类型ID错误：", err)
			return
		}
		artiType[i].ArticleCount = num//go特有的：不能通过“v.属性”来修改结构体，但是可以通过“对象[索引]”的方式修改
	}


	c.Data["articles"] = articles
	c.Data["artiType"] = artiType

	c.TplName = "content.html"
}

//文章删除操作
func (c *MainController) Delete() {
	id, err := c.GetInt("id")//前端通过使用url?id=xxx的方式向后端返回id，此处使用c.GetInt("id")的方式获取前端的id值
	if err != nil {
		beego.Info("获取文章ID失败：", err)
		return
	}
	err = models.DeleteArticle(id)
	if err != nil {
		beego.Info("文章删除失败：", err)
		return
	}
	beego.Info("文章删除成功！")
}

//退出登录
func (c *MainController) Exit() {
	c.DelSession("username")//删除session
	c.Redirect("/", 302)
}

//显示聊天页面
func (c *MainController) ShowChat() {
	c.TplName = "chat.html"
}
//处理聊天内容
func (c *MainController) Chat() {
	msg := c.GetString("message")
	if msg == "" {
		beego.Info("输入消息为空！")
		return
	}

	turingMsg, err := turing.Robots("your turing apikey", turing.ReqType(1), msg)
	if err != nil {
		beego.Info("图灵机器人解析错误：", err)
		return
	}

	c.Ctx.WriteString(msg+"**###**"+turingMsg.(string))
}

//显示编辑修改页
func (c *MainController) ShowRevise() {
	//1、获取文章ID
	artiId, err := c.GetInt("rid")
	//beego.Info("artiId-->", artiId)
	if err != nil {
		beego.Info("文章ID获取失败：", err)
		return
	}
	//2、根据文章ID查询文章标题、类别、内容、图片
	o := orm.NewOrm()
	arti := models.Article{Id:artiId}
	_, err = o.QueryTable("Article").Filter("Id", artiId).RelatedSel("ArticleType").All(&arti)
	if err != nil {
		beego.Info("查询文章出错：", err)
		return
	}

	c.Data["article"] = arti
	c.TplName = "revise.html"
}
//修改编辑操作
func (c *MainController) Revise() {
	artiId, err := c.GetInt("artiId")
	if err != nil {
		beego.Info("文章ID获取失败：", err)
		return
	}
	artiName := c.GetString("artiname")
	artiContent := c.GetString("content")
	artiImg := c.GetString("imgname")
	newTime := time.Now()

	if artiName == "" || artiContent == "" {
		beego.Info("非法输入！")
		return
	}
	resp, err := models.ArticleRevise(artiId, artiName, artiContent, newTime, artiImg)
	if resp == 0 || err != nil {
		beego.Info("文章修改失败：", err)
		c.Ctx.WriteString("failed")
		return
	}
	c.Ctx.WriteString("success")
}

//分类文章显示
func (c *MainController) ShowClassify() {
	o := orm.NewOrm()
	//1、显示侧边分类
	artiType, err := models.GetArtiType()
	if err != nil {
		beego.Info("类型查找失败：", err)
		return
	}
	//2、获取分类Id
	typeId, err := c.GetInt("tid")
	if err != nil {
		beego.Info("分类ID获取失败：", err)
		return
	}
	//3、根据分类Id查询该类下的文章
	article, err := models.GetTypeArti(typeId)
	if err != nil || article == nil {
		beego.Info("分类文章获取失败：", err)
		return
	}
	//在分类下显示该类数量
	for i, v:= range artiType {
		num, err := o.QueryTable("Article").Filter("ArticleType__id", v.Id).RelatedSel().Count()
		if err != nil {
			beego.Info("读取文章类型ID错误：", err)
			return
		}
		artiType[i].ArticleCount = num//go特有的：不能通过“v.属性”来修改结构体，但是可以通过“对象[索引]”的方式修改
	}
	//4、将查询到的文章发送至前端
	c.Data["articles"] = article
	c.Data["artiType"] = artiType
	c.TplName = "classify.html"
}

//图片欣赏页展示
func (c *MainController) Gallery() {
	o := orm.NewOrm()
	arti := []models.Article{}
	//SQL查询：select * from article where img is not null and img!=""
	_, err := o.QueryTable("Article").Filter("Img__isnull", false).Exclude("Img__exact", "").All(&arti)
	if err != nil {
		beego.Info("图片获取失败：", err)
		return
	}

	c.Data["ArtiImg"] = arti
	c.TplName = "gallery.html"
}

//文件下载
func (c *MainController) Download() {
	img := c.GetString("picId")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))//获取项目路径
	if err != nil {
		beego.Info("项目路径获取失败：", err)
		return
	}
	imgPath := dir + "/static/userImg/" + img
	c.Ctx.Output.Download(imgPath, img)//第一个参数为下载文件的全路径，第二个参数是保存的文件名（可选）
}

//显示404出错界面
func (c *MainController) ShowFalse()  {
	c.TplName = "404-page.html"
}
