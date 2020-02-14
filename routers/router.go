package routers

import (
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/controllers"
)

func init() {
	docs := beego.NewNamespace("docs",
		beego.NSInclude(&controllers.DocumentsController{}),
	)

	admin := beego.NewNamespace("login",
		beego.NSInclude(&controllers.LoginController{}),
	)

	beego.Router("/",&controllers.PageController{},"get:GetIndex")
	beego.Router("/background",&controllers.PageController{},"get:GetBackground")
	beego.AddNamespace(admin, docs)
}
