package main

import (
	_ "myBlog/models"
	_ "myBlog/routers"

	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//default 数据库别名，默认default
	//force true, 强制建表
	//verbose true 是否打印建表sql过程
	orm.RunSyncdb("default", false, false)

	beego.SetStaticPath("/static", "static")
	beego.BConfig.WebConfig.ViewsPath = "views"
	//session on
	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.WebConfig.Session.SessionName = "beegosession"

	beego.Run()
}
