package controllers

import (
	"Service_Monitor/models"
	"Service_Monitor/ssh_connector"
	"github.com/astaxie/beego/context"
	"os"
)

func UseSSH(services []models.Service) []models.Service {
	var ctx *context.Context
	//获取当前用户名和密码
	username := ctx.Input.Session("username")
	pwd := ctx.Input.Session("password")
	//fmt.Println(username1.(string)+":"+pwd1.(string))

	//username:= "root"
	//pwd := "1qaz@WSX"

	for i, v := range services {
		status := SSH(username.(string), pwd.(string), v.Ip, v.Name)
		if status == "0" {
			services[i].Status = "down"
		}else {
			services[i].Status = "up"
		}
	}
	return services
}

func SSH(user, password, ip, serviceName string) string {
	session, err := ssh_connector.Connect(user, password, ip, 22)
	if err != nil {
		return ""
	}
	defer session.Close()

	//session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	//ip_182.61.33.89 username:root passwd:1qaz@WSX
	//session.Run("netstat -lntup|grep nginx |wc -l") //sh 命令路径
	s, _ := session.Output("netstat -lntup|grep " + serviceName +  " |wc -l")
	return string(s)
}