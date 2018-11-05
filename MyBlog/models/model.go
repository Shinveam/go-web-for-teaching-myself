package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int
	UserName string
	Pwd string
}

func init()  {
	orm.Debug = true//开发环境下建议开启调试
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/myblog?charset=utf8")

	orm.RegisterModel(new(User))
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
func RegisterUser(name string, pwd string) (maps []orm.Params, err error) {
	search, err := SearchUser(name, pwd)
	if search != nil && err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	if search == nil && err != nil{
		user := User{}
		user.UserName = name
		user.Pwd = pwd

		_, err := o.Insert(&user)
		if err != nil {
			return nil, err
		}
		return maps, err //此时的err应为nil
	}

	return nil, err
}

