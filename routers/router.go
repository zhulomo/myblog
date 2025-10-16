package routers

import (
	"context"
	"myBlog/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.InsertFilter("/article/*", beego.BeforeRouter, LoginFilter)
	beego.Router("/article/list", &controllers.ArticleListController{})
	beego.Router("/api/articles", &controllers.ArticleListController{}, "get:List")
	beego.Router("/article/detail/:id", &controllers.ArticleDetailController{}, "get:ShowArticle")
	beego.Router("/article/add", &controllers.ArticleController{})
	beego.Router("/article/update", &controllers.ArticleController{})

}

var LoginFilter = func(ctx *context.Context) {
	username := ctx.Input().Session("username")
	if username == nil {
		ctx.Redirect(302, "/login")
	}
}
