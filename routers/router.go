package routers

import (
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/controllers"
)

var APIs = make(map[string]func())

func init() {
	beego.Router("/", &controllers.HomePageController{})
	beego.ErrorController(&controllers.ErrorController{})

	admin := beego.NewNamespace("admin",
		beego.NSCond(controllers.CheckToken),
		//beego.NSNamespace("ssh",
		//	beego.NSInclude(&controllers.SSHController{})),
		beego.NSNamespace("profile",
			beego.NSInclude(&controllers.ProfileController{})),

		beego.NSNamespace("login",
			beego.NSInclude(&controllers.LoginController{})),

		beego.NSNamespace("api",
			beego.NSInclude(&controllers.APIController{})), //write manager
	)
	beego.AddNamespace(admin)
	//beego.ErrorHandler("404" ,controllers.PageNotFound)
	//beego.Router("/all/:key", &CMSController{}, "get:AllBlock")
}
