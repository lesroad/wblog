package models

import (
	"bytes"
	"encoding/gob"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id              int
	Title           string
	Created         time.Time `orm:"auto_now;null"`
	Views           int       `orm:"index"`
	TopicTime       time.Time `orm:"auto_now;null"`
	TopicCount      int
	TopicLastUserId int
}

type Topic struct {
	Id              int
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Category        string
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int       `orm:"index"`
	Author          string
	Labels          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int
	ReplyLastUserId int
}

type Comment struct {
	Id      int
	Tid     int
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:1256560.@tcp(127.0.0.1:3306)/wuwen?charset=utf8")
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RunSyncdb("default", false, true)
}

func GetAllTopics(cate string, label string, isDesc bool) (topics []*Topic) {
	topics = make([]*Topic, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("Topic")
	if len(cate) > 0 {
		qs = qs.Filter("category", cate)
	}
	if len(label) > 0 {
		qs = qs.Filter("labels__contains", "$"+label+"#") //字段__包含
	}
	if isDesc {
		qs.OrderBy("-Updated").All(&topics)
	} else {
		qs.OrderBy("Updated").All(&topics)
	}
	return topics
}

func GetAllCategories() *[]Category {
	cates := make([]Category, 0)
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		beego.Info("redis启动错误err：", err)
	}
	defer conn.Close()
	rel, _ := redis.Bytes(conn.Do("get", "cates"))
	dec := gob.NewDecoder(bytes.NewReader(rel))
	dec.Decode(&cates)
	// fmt.Printf("从redis中读的 = %v", cates)

	if len(cates) == 0 {
		o := orm.NewOrm()
		qs := o.QueryTable("category")
		qs.All(&cates)

		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(cates)
		conn.Do("set", "cates", buf.Bytes())
		// fmt.Printf("这是在mysql中读的")
	}
	return &cates
}

func DeleteCategory(id string) {
	o := orm.NewOrm()
	cid, _ := strconv.Atoi(id)
	one := &Category{Id: cid}
	o.Read(one)
	_, err := o.Delete(one)
	if err != nil {
		beego.Info("err :", err)
	}
}

func ModifyCateNum(cate string, num int) {
	Cate := &Category{}
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	qs.Filter("title", cate).One(Cate)
	Cate.TopicCount = Cate.TopicCount + num
	o.Update(Cate)
}

func AddTopic(title, category, label, content, filename string) {
	//处理标签  得到格式 $I'm#$lesroad#
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	topic := Topic{
		Title:      title,
		Content:    content,
		Category:   category,
		Labels:     label,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
		Attachment: filename,
	}
	o.Insert(&topic)
	//对应的分类+1
	ModifyCateNum(category, 1)
}

func ModifyTopic(ttid, title, category, label, content, filename string) {
	//处理标签  得到格式 $I'm#$lesroad#
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	o := orm.NewOrm()
	tid, _ := strconv.Atoi(ttid)
	topic := &Topic{Id: tid}
	var oldCate, oldAttach string //保存旧的分类和旧的附件名称
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Category = category
		topic.Labels = label
		topic.Content = content
		topic.Updated = time.Now()
		topic.Attachment = filename
		o.Update(topic)
	}
	//更新分类统计
	if oldCate != category {
		ModifyCateNum(oldCate, -1)
		ModifyCateNum(category, +1)
	}
	//删除旧附件
	if len(oldAttach) > 0 {
		os.Remove("./upload file/" + oldAttach)
	}
}

func GetTopic(ttid string) *Topic {
	o := orm.NewOrm()
	tid, _ := strconv.Atoi(ttid)
	topic := &Topic{}
	o.QueryTable("topic").Filter("id", tid).One(topic)
	topic.Views++
	o.Update(topic)

	// #换成空格再把$换成空格
	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", " ", -1)
	return topic
}

func DeleteTopic(id string) {
	o := orm.NewOrm()
	tid, _ := strconv.Atoi(id)
	one := &Topic{Id: tid}
	o.Read(one)
	var oldcate = one.Category
	_, err := o.Delete(one)
	if err != nil {
		beego.Info("err :", err)
	}
	ModifyCateNum(oldcate, -1)
}

func ModifyTopicByReply(tid, num int) {
	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	o.Read(topic)
	topic.ReplyTime = time.Now()
	topic.ReplyCount = topic.ReplyCount + num
	beego.Info("id:评论：", tid, topic.ReplyCount)
	o.Update(topic)
}

func AddReply(tid, nickname, content string) {
	tidnum, _ := strconv.Atoi(tid)
	reply := &Comment{
		Tid:     tidnum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	o.Insert(reply)

	beego.Info("评论的tid：", tidnum)
	//添加评论后修改文章属性
	ModifyTopicByReply(tidnum, 1)
}

func GetAllReplies(tid string) (reply []*Comment) {
	tidnum, _ := strconv.Atoi(tid)
	replies := make([]*Comment, 0)

	o := orm.NewOrm()
	o.QueryTable("comment").Filter("tid", tidnum).All(&replies) //这里一定加取址
	return replies
}

func DeleteReply(rid string) {
	ridnum, _ := strconv.Atoi(rid)
	o := orm.NewOrm()
	one := &Comment{Id: ridnum}
	o.Read(one)
	tid := one.Tid
	o.Delete(one)
	//删除评论后修改文章属性
	ModifyTopicByReply(tid, -1)
}
