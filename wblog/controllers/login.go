package controllers

import (
	"beego项目/wblog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	name := c.Ctx.GetCookie("uname") //自动登录
	c.Data["name"] = name
	pwd := c.Ctx.GetCookie("upwd")
	c.Data["pwd"] = pwd

	c.TplName = "login.html"
}
func (c *LoginController) Post() {
	name := c.GetString("uname")
	pwd := c.GetString("upwd")
	if err := models.CheckNameAndPwd(name, pwd); err != nil {
		beego.Info("用户名密码错误！")
		c.TplName = "login.html"
		return
	}

	auto := c.GetString("autoLogin")
	if auto == "on" {
		c.Ctx.SetCookie("uname", name, 3600)
		c.Ctx.SetCookie("upwd", pwd, 3600)
	} else {
		c.Ctx.SetCookie("uname", "xjbx", -1)
		c.Ctx.SetCookie("upwd", "xjbx", -1)
	}
	c.SetSession("name", name) //设置session，提前设置conf
	c.Redirect("/", 302)
}

func (c *LoginController) Quit() {
	c.DelSession("name") //删除session
	c.Redirect("/", 302)
}
