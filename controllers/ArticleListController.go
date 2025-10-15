package controllers

import (
	"myBlog/models"

	"github.com/beego/beego/v2/server/web"
)

type ArticleListController struct {
	web.Controller
}

func (R *ArticleListController) Get() {
	R.TplName = "articlelist.html"
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
