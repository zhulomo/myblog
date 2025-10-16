package controllers

import (
	"fmt"
	"myBlog/models"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// 显示登录页面
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

// 登录 POST
func (c *LoginController) Post() {
	fmt.Printf("Request URL: %s, Method: %s\n", c.Ctx.Request.URL.Path, c.Ctx.Request.Method)
	fmt.Printf("Before SetSession, Controller: %+v\n", c.Controller)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Login Panic:", r)
			c.Data["json"] = JsonResponse{Code: 0, Message: "服务器错误"}
			c.ServeJSON()
		}
	}()

	username := c.GetString("username")
	password := c.GetString("password")

	if username == "" || password == "" {
		c.Data["json"] = JsonResponse{Code: 0, Message: "用户名或密码不能为空"}
		c.ServeJSON()
		return
	}

	pwd, err := models.GetUserPasswordByName(username)
	if err != nil {
		c.Data["json"] = JsonResponse{Code: 0, Message: "用户名不存在"}
		c.ServeJSON()
		return
	}

	if pwd == password {
		// ⚠️ 不要手动设置 CruSession
		c.SetSession("username", username)

		// 验证 Session 是否设置成功
		sessionUser := c.GetSession("username")
		fmt.Printf("Session username after set: %v\n", sessionUser)

		c.Data["json"] = JsonResponse{Code: 1, Message: "登录成功"}
		c.ServeJSON()
	} else {
		c.Data["json"] = JsonResponse{Code: 0, Message: "密码错误"}
		c.ServeJSON()
	}
}
