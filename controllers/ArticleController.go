package controllers

import (
	"myBlog/models"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	web.Controller
}

func (R *ArticleController) Get() {
	R.TplName = "article_add.html"
}

func (R *ArticleController) Post() {
	// 	type Article struct {
	// 	Id         int       `orm:"auto"`
	// 	Title      string    `orm:"column(title);size(64)"`
	// 	Abstract   string    `orm:"column(abstract);size(64)"`
	// 	Content    string    `orm:"column(content)type(text)"`
	// 	Author     string    `orm:"column(author);size(64)"`
	// 	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	// }
	title := R.GetString("title")
	abstract := R.GetString("abstract")
	content := R.GetString("content")
	author := R.GetString("author")

	// user := models.Users{
	// 	Username:   username,
	// 	Password:   password,
	// 	Status:     0,
	// 	CreateTime: time.Now(),
	// }
	article := models.Article{
		Title:      title,
		Abstract:   abstract,
		Content:    content,
		Author:     author,
		CreateTime: time.Now(),
	}

	_, err := models.ArticleInsert(&article)

	if err != nil {
		R.Ctx.WriteString("出错")
		return
	}

	R.Ctx.WriteString("发布成功")
}
