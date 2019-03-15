package controllers

import (
	"beego项目/wblog/models"
	"strings"
)

type TopicController struct {
	BaseController
}

func (c *TopicController) Get() {
	topics := models.GetAllTopics(c.GetString("cate"), "", false) //分类， label, 排序
	c.Data["topics"] = topics
	c.TplName = "topic.html"
}

func (c *TopicController) Add() {
	if err := c.CheckLogin(); err != nil {
		return
	}
	c.TplName = "add.html"
}

func (c *TopicController) Post() {
	tid := c.GetString("tid")
	title := c.GetString("title")
	content := c.GetString("content")
	category := c.GetString("category")
	label := c.GetString("label")
	//获取附件
	_, h, _ := c.GetFile("file")
	var filename string
	if h != nil {
		//上传了附件
		filename = h.Filename
		c.SaveToFile("file", "./upload file/"+filename)
	}
	if len(tid) == 0 {
		models.AddTopic(title, category, label, content, filename)

	} else {
		models.ModifyTopic(tid, title, category, label, content, filename)
	}

	c.Redirect("/topic", 302)
}

func (c *TopicController) View() {
	topic := models.GetTopic(c.Ctx.Input.Param("0"))
	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
	c.TplName = "content.html"

	//取出评论
	replies := models.GetAllReplies(c.Ctx.Input.Param("0"))
	c.Data["Replies"] = replies

	c.Data["Labels"] = strings.Split(topic.Labels, " ")
}

func (c *TopicController) Modify() {
	if err := c.CheckLogin(); err != nil {
		return
	}
	tid := c.GetString("tid")
	topic := models.GetTopic(tid)
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
	c.TplName = "modify.html"
}

func (c *TopicController) Delete() {
	if err := c.CheckLogin(); err != nil {
		return
	}
	models.DeleteTopic(c.GetString("tid"))
	c.Redirect("/topic", 302)
}
