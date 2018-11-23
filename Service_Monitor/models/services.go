package models

import (
	"github.com/astaxie/beego/orm"
	"os/exec"
)

//查找服务
func SearchService(id int) (Services []Service, err error) {
	o := orm.NewOrm()

	if id == 0 {
		_, err = o.QueryTable("Service").All(&Services)
		if err != nil {
			return nil, err
		}
		return Services, err
	} else{

		_, err = o.QueryTable("Service").Filter("Id", id).All(&Services)
		if err != nil {
			return nil, err
		}
		return Services, err
	}
	return nil, err
}
//添加服务
func AddService(ip string, name string, email string) (int, error) {
	o := orm.NewOrm()
	services := Service{}

	services.Ip = ip
	services.Name = name
	services.Email = email

	_, err := o.Insert(&services)
	if err != nil {
		return 0, err
	}

	return 1, err
}
//删除服务
func DeleteService(id int) (int, error) {
	o := orm.NewOrm()
	services := Service{Id:id}
	_, err := o.Delete(&services)
	if err != nil {
		return 0, err
	}
	return 1, err
}
//修改服务
func ReviseService(id int, ip string, name string, email string) (int,  error) {
	o := orm.NewOrm()
	services := Service{Id:id}
	err := o.Read(&services)
	if err != nil {
		return 0, err
	}else {
		services.Ip = ip
		services.Name = name
		services.Email = email
		_, err = o.Update(&services)
		if err != nil {
			return 0, err
		}
		return 1, err
	}
	return 0, err
}

//查询服务状态--将函数名该为init
func ServiceStatus() {
	//1、查询所有服务
	o := orm.NewOrm()
	var Services []Service
	_, err := o.QueryTable("Service").All(&Services)
	if err != nil {
		return
	}
	//2、发送shell命令
	for i, v := range Services {
		serviceName := v.Name
		cmd := exec.Command("netstat -lntup|grep " + serviceName +" |wc -l")
		num, err := cmd.Output()
		status := string(num)
		if  err != nil {
			return
		}
		//return err
		//将状态保存在数据库
		if status == "0" {
			Services[i].Status = "down"
			_, err = o.Update(&Services[i])
			if err != nil {
				return
			}
		}else {
			Services[i].Status = "up"
			_, err = o.Update(&Services[i])
			if err != nil {
				return
			}
		}
	}
}