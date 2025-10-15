package controllers

import (
	"myBlog/models"

	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func (R *LoginController) Get() {
	R.TplName = "login.html"
}

func (R *LoginController) Post() {
	username := R.GetString("username")
	password := R.GetString("password")

	userpassword, err := models.GetUserPasswordByName(username)
	if err != nil {
		R.Data["json"] = JsonResponse{Code: 0, Message: "用户名不存在"}
		R.ServeJSON()
		return
	}

	if password == userpassword {
		R.Data["json"] = JsonResponse{Code: 1, Message: "登录成功"}
		R.ServeJSON()

	} else {
		R.Data["json"] = JsonResponse{Code: 0, Message: "密码错误"}
		R.ServeJSON()
	}

}
