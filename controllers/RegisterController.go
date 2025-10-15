package controllers

import (
	"myBlog/models"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	web.Controller
}
type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (R *RegisterController) Get() {
	R.TplName = "register.html"
}

func (R *RegisterController) Post() {
	username := R.GetString("username")
	password := R.GetString("password")
	repassword := R.GetString("repassword")

	if password != repassword {
		R.Data["json"] = JsonResponse{Code: 0, Message: "两次输入的密码不一致"}
		R.ServeJSON()
		return
	}

	existuser, err := models.GetUserByName(username)
	if err != nil {
		R.Data["json"] = JsonResponse{Code: 0, Message: "查询失败"}
		R.ServeJSON()
		return
	}

	if existuser != nil {
		R.Data["json"] = JsonResponse{Code: 0, Message: "用户已存在"}
		R.ServeJSON()
		return
	}
	user := models.Users{
		Username:   username,
		Password:   password,
		Status:     0,
		CreateTime: time.Now(),
	}

	_, err = models.InsertUser(&user)
	if err != nil {
		R.Data["json"] = JsonResponse{Code: 0, Message: "注册失败"}
		R.ServeJSON()
	}
	R.Data["json"] = JsonResponse{Code: 1, Message: "注册成功"}
	R.ServeJSON()

}
