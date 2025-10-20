package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

// 判断是否登录
func (R *BaseController) Prepare() {
	R.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	R.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	R.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	R.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")

	username := R.GetSession("username")
	fmt.Println("loginuser---->", username)
	if username != nil {
		R.IsLogin = true
		R.Loginuser = username
	} else {
		R.IsLogin = false
	}
	R.Data["IsLogin"] = R.IsLogin
}
func (c *BaseController) Options() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Ctx.WriteString("ok") // 注意：必须写，不然 header 可能不生效
}
