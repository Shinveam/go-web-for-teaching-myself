package main

import (
	_ "article/routers"
	_ "article/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

