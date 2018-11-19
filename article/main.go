package main

import (
	_ "article/routers"
	_ "article/models"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	beego.AddFuncMap("PreviousPage", HandlePreviousPage)//给前端的PreviousPage和后端的HandlePreviousPage函数做映射
	beego.AddFuncMap("NextPage", HandleNextPage)//原理同上
	beego.Run()
}

//上一页处理
func HandlePreviousPage(page int) (pageindex string) {
	pageindex = strconv.Itoa(page-1)
	return pageindex
}
//下一页处理
func HandleNextPage(page int) string {
	pageindex := strconv.Itoa(page+1)
	return pageindex
}

