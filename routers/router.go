package routers

import (
	"myBlog/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/articlelist", &controllers.ArticleListController{})
	beego.Router("/api/articles", &controllers.ArticleListController{}, "get:List")
}
