package main

import (
	"github.com/astaxie/beego"
	_ "github.com/painterQ/painterBlog/routers"
)

//go:generate  go generate ./vue
func main() {
	beego.DelStaticPath("static")
	beego.SetStaticPath("/", "./static")
	beego.Run()
}
