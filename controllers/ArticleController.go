package controllers

import (
	"encoding/json"
	"fmt"
	"myBlog/models"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

type ArticleController struct {
	BaseController
}

func (R *ArticleController) Get() {
	id, _ := R.GetInt("id")
	loginuser := R.GetSession("username")
	fmt.Println("updateloginuser:", loginuser)
	//if id == 0 {
	//R.TplName = "article/add.html"
	//} else {
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)
	if err != nil {
		R.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "文章不存在",
		}
		R.ServeJSON()
		return
	}

	if article.Author == loginuser {
		article := models.Article{Id: id}
		o := orm.NewOrm()
		_ = o.Read(&article)
		R.Data["json"] = map[string]interface{}{
			"code":    1,
			"article": article,
		}
		R.ServeJSON()
		return
		//R.TplName = "article/update.html"
	} else {
		R.Data["json"] = map[string]interface{}{"code": 0, "msg": "您无权限修改文章"}
		R.ServeJSON()
	}

	//}
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
	// title := R.GetString("title")
	// abstract := R.GetString("abstract")
	// content := R.GetString("content")
	var req struct {
		Title    string `json:"title"`
		Abstract string `json:"absract"`
		Content  string `json:"content"`
	}
	body := R.Ctx.Input.RequestBody
	json.Unmarshal(body, &req)

	loginuser := R.GetSession("username")
	author, ok := loginuser.(string)
	if !ok {
		R.Data["json"] = map[string]interface{}{"code": 0, "msg": "未登录或session失效"}
		R.ServeJSON()
		return
	}

	// user := models.Users{
	// 	Username:   username,
	// 	Password:   password,
	// 	Status:     0,
	// 	CreateTime: time.Now(),
	// }
	var article models.Article
	if id == 0 {
		article = models.Article{
			Title:      req.Title,
			Abstract:   req.Abstract,
			Content:    req.Content,
			Author:     author,
			CreateTime: time.Now(),
		}

		_, err := models.ArticleInsert(&article)

		if err != nil {
			R.Data["json"] = map[string]interface{}{"code": 0, "msg": "出错"}
		} else {
			R.Data["json"] = map[string]interface{}{"code": 1, "msg": "发布文章成功"}
		}
		R.ServeJSON()

	} else {
		article.Id = id
		article.Title = req.Title
		article.Abstract = req.Abstract
		article.Author = author
		article.Content = req.Content

		// o := orm.NewOrm
		// error := o.Read(&article)
		num, err := models.ArticleUpdate(&article)
		if err != nil {
			R.Data["json"] = map[string]interface{}{"code": 0, "msg": "更新出错"}
			return
		}
		if num != 0 {
			R.Data["json"] = map[string]interface{}{"code": 1, "msg": "更新成功"}
		}
		R.ServeJSON()

	}

}
func (R *ArticleController) DeleteById() {
	//id, _ := R.GetInt("id")
	//author := R.GetString("author")
	loginuser := R.GetSession("username")
	var req struct {
		Id int `json:"id"`
	}
	json.Unmarshal(R.Ctx.Input.RequestBody, &req)
	id := req.Id
	article, err := models.GetArticleById(id)
	if err != nil {
		R.Data["json"] = map[string]interface{}{"code": 0, "mas": "文章不存在"}
		R.ServeJSON()
		return
	}
	author := article.Author
	if author == loginuser {

		num, err := models.ArticleDelete(id)
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
