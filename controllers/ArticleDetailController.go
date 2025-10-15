package controllers

import (
	"myBlog/models"

	"github.com/beego/beego/v2/server/web"
)

type ArticleDetailController struct {
	web.Controller
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
		R.Ctx.WriteString("发生奇怪的错误")
		return
	}

	if aritcle == nil {
		R.Ctx.WriteString("文章消失了")
		return
	}
	R.Data["Article"] = aritcle
	R.TplName = "article_detail.html"
}
