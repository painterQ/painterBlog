package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/controllers"
	"github.com/painterQ/painterBlog/models"
	_ "github.com/painterQ/painterBlog/routers"
	"path"
)

//go:generate  go generate ./vue
func main() {
	//err
	beego.ErrorHandler("403", controllers.Err403)
	beego.ErrorHandler("404", controllers.Err404)

	//logs
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)

	//filter
	tf := controllers.TokenFilter{}
	beego.InsertFilter(tf.GetPattern(), tf.GetPosition(), tf.GetFilter())

	beego.DelStaticPath("static")
	dbPath := beego.AppConfig.DefaultString(models.ConfigDBPath, models.DefaultDBPathConfig)
	beego.SetStaticPath("/image",path.Join(dbPath, models.ImagePath))
	beego.SetStaticPath("/static", "./static")
	beego.Run()
}
