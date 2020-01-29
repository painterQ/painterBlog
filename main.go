package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/painterQ/painterBlog/routers"
)

//go:generate  go generate ./vue
func main() {
	//logs
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)

	beego.DelStaticPath("static")
	beego.SetStaticPath("/static", "./static")
	beego.Run()
}
