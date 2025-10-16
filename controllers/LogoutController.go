package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	beego.Controller
}

func (R *LogoutController) Get() {
	R.DestroySession()
	R.Redirect("/login", 302)
}
