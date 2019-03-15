package routers

import (
	"beego项目/wblog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})

	beego.Router("/alogin", &controllers.LoginController{})
	beego.Router("/quit", &controllers.LoginController{}, "get:Quit")
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/delete", &controllers.CategoryController{}, "get:Delete")

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

}
