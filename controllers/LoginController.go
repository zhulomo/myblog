package controllers

import (
	"encoding/json"
	"fmt"
	"myBlog/models"
)

type LoginController struct {
	BaseController
}

// 显示登录页面
// func (c *LoginController) Get() {
// 	// c.TplName = "login.html"
// }

// 登录 POST
func (c *LoginController) Post() {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body := c.Ctx.Input.RequestBody

	json.Unmarshal(body, &req)
	password, err := models.GetUserPasswordByName(req.Username)
	if err != nil || password != req.Password {
		c.Data["json"] = map[string]interface{}{"code": 0, "msg": "用户名或密码错误"}
	} else {
		c.SetSession("username", req.Username)
		fmt.Println("usersession---->", req.Username)

		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "登录成功"}
	}
	c.ServeJSON()
}

// 判断是否登录
func (R *BaseController) LoginVerify() {
	username := R.GetSession("username")
	fmt.Println("loginuser---->", username)
	if username != nil {
		R.Data["json"] = map[string]interface{}{
			"code":     1,
			"isLogin":  true,
			"username": username,
		}
		// R.IsLogin = true
		// R.Loginuser = username
	} else {
		R.Data["json"] = map[string]interface{}{
			"code":    0,
			"isLogin": false,
		}
	}
	R.ServeJSON()
}

// func (c *LoginController) Post() {
// 	fmt.Printf("Request URL: %s, Method: %s\n", c.Ctx.Request.URL.Path, c.Ctx.Request.Method)
// 	fmt.Printf("Before SetSession, Controller: %+v\n", c.Controller)
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Login Panic:", r)
// 			c.Data["json"] = JsonResponse{Code: 0, Message: "服务器错误"}
// 			c.ServeJSON()
// 		}
// 	}()

// 	username := c.GetString("username")
// 	password := c.GetString("password")

// 	if username == "" || password == "" {
// 		c.Data["json"] = JsonResponse{Code: 0, Message: "用户名或密码不能为空"}
// 		c.ServeJSON()
// 		return
// 	}

// 	pwd, err := models.GetUserPasswordByName(username)
// 	if err != nil {
// 		c.Data["json"] = JsonResponse{Code: 0, Message: "用户名不存在"}
// 		c.ServeJSON()
// 		return
// 	}

// 	if pwd == password {
// 		// ⚠️ 不要手动设置 CruSession
// 		c.SetSession("username", username)

// 		// 验证 Session 是否设置成功
// 		sessionUser := c.GetSession("username")
// 		fmt.Printf("Session username after set: %v\n", sessionUser)

// 		c.Data["json"] = JsonResponse{Code: 1, Message: "登录成功"}
// 		c.ServeJSON()
// 	} else {
// 		c.Data["json"] = JsonResponse{Code: 0, Message: "密码错误"}
// 		c.ServeJSON()
// 	}
// }
