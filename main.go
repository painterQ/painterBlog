package main

import (
	"github.com/astaxie/beego"
	_ "github.com/painterQ/painterBlog/routers"
)


//go:generate  cd ./vue && cnpm run build
func main() {
	beego.Run()
}
