package controllers

import (
	"myBlog/models"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

type ArticleController struct {
	BaseController
}

func (R *ArticleController) Get() {
	id, _ := R.GetInt("id")

	if id == 0 {
		R.TplName = "article/add.html"
	} else {
		article := models.Article{Id: id}
		o := orm.NewOrm()
		_ = o.Read(&article)
		R.Data["Article"] = article
		R.TplName = "article/update.html"
	}
}

func (R *ArticleController) Post() {
	// 	type Article struct {
	// 	Id         int       `orm:"auto"`
	// 	Title      string    `orm:"column(title);size(64)"`
	// 	Abstract   string    `orm:"column(abstract);size(64)"`
	// 	Content    string    `orm:"column(content)type(text)"`
	// 	Author     string    `orm:"column(author);size(64)"`
	// 	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	//
	id, _ := R.GetInt("id")
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
	var article models.Article
	if id == 0 {
		article = models.Article{
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
	} else {
		article.Id = id
		article.Title = title
		article.Abstract = abstract
		article.Author = author
		article.Content = content

		// o := orm.NewOrm
		// error := o.Read(&article)
		num, err := models.ArticleUpdate(&article)
		if err != nil {
			R.Ctx.WriteString("更新出错")
			return
		}
		if num != 0 {
			R.Ctx.WriteString("更新成功")
		}
	}

}
func (R *ArticleController) DeleteById() {
	id, _ := R.GetInt("id")
	author := R.GetString("author")
	loginuser := R.Loginuser
	if author == loginuser {
		article := models.Article{
			Id: id,
		}
		num, err := models.ArticleDelete(&article)
		if err != nil {
			R.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除失败"}
			R.ServeJSON()
			return
		}
		if num != 0 {
			R.Data["json"] = map[string]interface{}{"code": 1, "msg": "删除成功"}
			R.ServeJSON()

		} else {
			R.Data["json"] = map[string]interface{}{"code": 0, "msg": "文章不存在或删除失败"}
		}
		R.ServeJSON()
	} else {
		R.Data["json"] = map[string]interface{}{"code": 0, "msg": "您无权限删除文章"}
		R.ServeJSON()
	}

}

// func (R *ArticleController) UpdateArticle() {
// 	id, _ := R.GetInt("Id")
// 	title := R.GetString("title")
// 	abstract := R.GetString("abstract")
// 	content := R.GetString("content")
// 	author := R.GetString("author")

// 	article := models.Article{
// 		Id:       id,
// 		Title:    title,
// 		Abstract: abstract,
// 		Content:  content,
// 		Author:   author,
// 	}
// 	// o := orm.NewOrm
// 	// error := o.Read(&article)

// 	num, err := models.ArticleUpdate(&article)
// 	if err != nil {
// 		R.Ctx.WriteString("更新出错")
// 		return
// 	}
// 	if num != 0 {
// 		R.Ctx.WriteString("更新成功")
// 	}

// }
