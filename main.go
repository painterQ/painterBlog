package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/controllers"
	"github.com/painterQ/painterBlog/models"
	_ "github.com/painterQ/painterBlog/routers"
	"os"
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

	beego.AddTemplateExt(".html")
	pwd,_ := os.Getwd()
	beego.BConfig.WebConfig.ViewsPath =  path.Join(pwd,"static")

	beego.DelStaticPath("static")
	dbPath := beego.AppConfig.DefaultString(models.ConfigDBPath, models.DefaultDBPathConfig)
	beego.SetStaticPath("/image",path.Join(dbPath, models.ImagePath))
	beego.SetStaticPath("/public", "./static/public")
	beego.SetStaticPath("/tinymce", "./static/tinymce")

	beego.Run()
}
