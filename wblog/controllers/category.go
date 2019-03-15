package controllers

import (
	"beego项目/wblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Get() {
	cates := models.GetAllCategories()
	c.Data["Categories"] = cates
	c.Data["ifexist"] = 0
	c.Data["ifdelete"] = false
	c.TplName = "category.html"
}

func (c *CategoryController) Post() {
	name := c.GetString("name")
	o := orm.NewOrm()
	one := models.Category{Title: name}
	err := o.Read(&one, "Title")
	if err == nil {
		c.Data["ifexist"] = 1

	} else {
		_, err := o.Insert(&one)
		if err != nil {
			beego.Info("插入错误：", err)
		}
		c.Data["ifexist"] = 2
	}
	var cates []models.Category
	o.QueryTable("Category").All(&cates)
	c.Data["Categories"] = cates
	c.Data["ifdelete"] = false
	c.TplName = "category.html"
}

func (c *CategoryController) Delete() {
	id := c.GetString("id")
	models.DeleteCategory(id)
	c.Redirect("/category", 302)
}
