package controllers

import (
	//"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	//IsLogin   bool
	//Loginuser interface{}
}

// 判断是否登录
// func (R *BaseController) Prepare() {
// 	username := R.GetSession("username")
// 	fmt.Println("loginuser---->", username)
// 	if username != nil {
// 		R.IsLogin = true
// 		R.Loginuser = username
// 	} else {
// 		R.IsLogin = false
// 	}
// 	R.Data["IsLogin"] = R.IsLogin
// }
