package main

import (
	_ "MyBlog/routers"
	_ "MyBlog/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

