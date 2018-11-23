package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int     `orm:"pk"`
	Name string
	Pwd  string
}

type Service struct {
	Id     int `orm:"pk;auto"`
	Ip     string
	Name   string
	Email  string
	Status string
}

func init()  {
	orm.Debug = true//开发环境下建议开启调试
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/service_monitor?charset=utf8")

	orm.RegisterModel(new(User), new(Service))
	orm.RunSyncdb("default", false, true)
}