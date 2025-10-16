package controllers

import (
	"myBlog/models"

	beego "github.com/beego/beego/v2/server/web"
)

type ArticleListController struct {
	beego.Controller
}

func (R *ArticleListController) Get() {
	// username := R.GetSession("username")
	// if username == nil {
	// 	R.Redirect("/login", 302)
	// 	return
	// }
	// R.Data["username"] = username
	R.TplName = "article/list.html"
}

func (R *ArticleListController) List() {
	articles, err := models.GetAllArticles()

	if err != nil {
		R.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "获取文章失败",
		}
		R.ServeJSON()
		return
	}
	R.Data["json"] = map[string]interface{}{
		"code": 1,
		"data": articles,
	}
	R.ServeJSON()

}
