package controllers

import (
	"errors"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//判断显示管理员登录/退出
func (c *BaseController) Prepare() {
	c.Data["Path"] = c.Ctx.Request.RequestURI
	beego.Info("当前点击路径：", c.Ctx.Request.RequestURI)
	//检查session判断是否显示管理员登录
	name := c.GetSession("name")
	if name != nil {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsLogin"] = false
	}
}

func (c *BaseController) CheckLogin() (err error) {
	name := c.GetSession("name")
	if name == nil { //没有登录不能操作
		c.Redirect("/alogin", 302)
		err = errors.New("no session!")
	}
	return err
}
