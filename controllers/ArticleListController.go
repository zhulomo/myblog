package controllers

import (
	"fmt"
	"myBlog/models"
)

type ArticleListController struct {
	BaseController
}

// func (R *ArticleListController) Get() {
// 	// username := R.GetSession("username")
// 	// if username == nil {
// 	// 	R.Redirect("/login", 302)
// 	// 	return
// 	// }
// 	// R.Data["username"] = username
// 	//R.TplName = "article/list.html"
// }

func (R *ArticleListController) List() {

	page, _ := R.GetInt("page", 1)
	pageSize, _ := R.GetInt("pageSize", 5)
	articles, total, err := models.GetArticleByPage(page, pageSize)

	//articles, err := models.GetAllArticles()
	username := R.GetSession("username")
	fmt.Println("session is", username)

	if err != nil {
		R.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "获取文章失败",
		}
		R.ServeJSON()
		return
	}
	R.Data["json"] = map[string]interface{}{
		"code":     1,
		"data":     articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
	R.ServeJSON()

}
