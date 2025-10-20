package controllers

import (
	"encoding/json"
	"myBlog/models"
	"time"
)

type RegisterController struct {
	BaseController
}
type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (R *RegisterController) Get() {
	//R.TplName = "register.html"
}

func (R *RegisterController) Post() {
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Repassword string `json:"repassword"`
	}

	body := R.Ctx.Input.RequestBody

	json.Unmarshal(body, &req)

	if req.Password != req.Repassword {
		R.Data["json"] = JsonResponse{Code: 0, Message: "两次输入的密码不一致"}
		R.ServeJSON()
		return
	}

	existuser, err := models.GetUserByName(req.Username)
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
		Username:   req.Username,
		Password:   req.Password,
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
