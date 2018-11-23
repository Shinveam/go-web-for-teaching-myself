package models

import "github.com/astaxie/beego/orm"

//用户查找
func SearchUser(name string, pwd string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw("select * from user where name = ? and pwd = ?", name, pwd).Values(&maps)
	if err != nil || num == 0 {
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
		user.Name = name
		user.Pwd = pwd

		_, err := o.Insert(&user)
		if err != nil {
			return 0, err
		}
		return 1, err //此时的err应为nil
	}

	return 0, err
}