package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	UserName string
	Pwd string
}

type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string
	Content string
	PubTime string //`orm:"auto_now_add;type(datetime)"`
	Img string
	Count int
}

func init()  {
	orm.Debug = true//开发环境下建议开启调试
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/articles?charset=utf8")

	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}

//用户查找
func SearchUser(name string, pwd string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw("select * from user where user_name = ? and pwd = ?", name, pwd).Values(&maps)
	if err != nil && num == 0 {
		return nil, err
	}

	return maps, err
}

//用户注册
func RegisterUser(name string, pwd string) (ret int, err error) {
	search, err := SearchUser(name, pwd)
	if err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	if search == nil {
		user := User{}
		user.UserName = name
		user.Pwd = pwd

		_, err := o.Insert(&user)
		if err != nil {
			return 0, err
		}
		return 1, err //此时的err应为nil
	}

	return 0, err
}

//创建文章
func CreatArticle(artiname string, content string, timer time.Time) (ret int, err error) {
	o := orm.NewOrm()
	arti :=  Article{}

	arti.ArtiName = artiname
	arti.Content = content
	arti.PubTime = timer.Format("2006-01-02 15:04:05")

	_, err = o.Insert(&arti)
	if err != nil {
		return 0, err
	}
	return 1, err
}