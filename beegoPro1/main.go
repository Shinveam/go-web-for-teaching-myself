package main

import (
	_ "beegoPro1/routers"
	_ "beegoPro1/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

