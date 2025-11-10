package routers

import (
	"myBlog/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.HomeController{})

	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/logout", &controllers.LogoutController{})
	// beego.InsertFilter("/article/*", beego.BeforeRouter, LoginFilter)
	//beego.Router("/api/article/list", &controllers.ArticleListController{})
	beego.Router("/api/loginverify", &controllers.LoginController{}, "get:LoginVerify;options:Options")
	beego.Router("/api/articles", &controllers.ArticleListController{}, "get:List;options:Options")
	beego.Router("/api/article/detail/:id", &controllers.ArticleDetailController{}, "get:ShowArticle;options:Options")
	beego.Router("/api/article/add", &controllers.ArticleController{})
	beego.Router("/api/article/update", &controllers.ArticleController{})
	beego.Router("/api/article/delete", &controllers.ArticleController{}, "options:Options;post:DeleteById")

}

// var LoginFilter = func(ctx *context.Context) {
//     username := ctx.input
//     if username == nil {
//         ctx.Redirect(302, "/login")
//     }
// }
