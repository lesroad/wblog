package main

import (
	_ "beego项目/wblog/models"
	_ "beego项目/wblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	//获取文章附件的方式
	//静态文件映射 （url，文件夹相对路径）
	beego.SetStaticPath("/file", "upload file")
	beego.Run()
}
