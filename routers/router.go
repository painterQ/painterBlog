package routers

import (
	"github.com/painterQ/painterBlog/painterBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
