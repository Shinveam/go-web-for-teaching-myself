package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	)

type User struct {
	Id int
	Name string
	Pwd string
}

func init()  {
	//设置数据库基本信息(相当于连接数据库)
	// 第一个参数是数据库别名，一般用不到
	// 第三个参数是数据源，root是用户名，123456是密码，test1是数据库名称
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/beegotest?charset=utf8")
	//映射model数据,new（表名）创建表，后面可用“,”分隔开在new（表名）
	orm.RegisterModel(new(User))
	//生产表，第一个参数是使用的数据库名称
	// 第二个参数是是否强制更新，除非表建错，需要重新生成表，否则为false
	//第三个参数是是否在终端显示表的创建过程
	orm.RunSyncdb("default", false, true)
}
