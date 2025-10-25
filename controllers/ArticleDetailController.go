package controllers

import (
	"myBlog/models"
)

type ArticleDetailController struct {
	BaseController
}

//路由中指定了get跳转的ShowArticle方法，就不用再写通用的get方法了
// func (R *ArticleDetailController) Get() {
// 	R.TplName = "article_detail.html"
// }

func (R *ArticleDetailController) ShowArticle() {
	id, err := R.GetInt(":id")

	if err != nil {
		R.Ctx.WriteString("无效的文章 ID")
		return
	}

	aritcle, err := models.GetArticleById(id)

	if err != nil {
		R.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "获取失败",
		}
		R.ServeJSON()
		return
	}

	if aritcle == nil {
		R.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "文章消失了",
		}
		R.ServeJSON()
		return
	}
	R.Data["json"] = map[string]interface{}{
		"code":    1,
		"article": aritcle,
	}
	R.ServeJSON()

}
