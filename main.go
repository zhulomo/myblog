package main

import (
	_ "myBlog/models"
	_ "myBlog/routers"

	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.RunSyncdb("default", false, true)
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
