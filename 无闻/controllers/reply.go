package controllers

import (
	"beego项目/无闻/models"
)

type ReplyController struct {
	BaseController
}

func (c *ReplyController) Add() {
	tid := c.GetString("tid")
	nickname := c.GetString("nickname")
	content := c.GetString("content")
	models.AddReply(tid, nickname, content)
	c.Redirect("/topic/view/"+tid, 302)
}

func (c *ReplyController) Delete() {
	tid := c.GetString("tid")
	rid := c.GetString("rid")
	models.DeleteReply(rid)
	c.Redirect("/topic/view/"+tid, 302)
}