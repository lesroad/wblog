package controllers

import (
	"beego项目/无闻/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	cate := c.GetString("cate")
	label := c.GetString("label")
	topics := models.GetAllTopics(cate, label, true)
	c.Data["Topics"] = topics
	c.TplName = "home.html"

	categories := models.GetAllCategories()
	c.Data["Categories"] = categories
}
