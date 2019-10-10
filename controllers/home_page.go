package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models"
	"strconv"
)

type HomePageController struct {
	beego.Controller
}

func (hp *HomePageController)Get() {
	hp.Data = map[interface{}]interface{}{
		"BlogName": "BlogName",
	}
	hp.Data["Version"] = "1.0.1"//StaticVersion(c)
	hp.Data["Title"] = `home page`
	hp.Data["Description"] = "博客首页，" + `Ei.SubTitle`
	hp.Data["Path"] = hp.Ctx.Input.URL()
	hp.Data["CurrentPage"] = "blog-home"
	hp.Data["Manager"] = models.ManagerSingleCase
	hp.Data["Blog"] = models.BlogSingleCase
	hp.Data["Qiniu"] = models.CDNSingleCase
	pnStr := hp.Input().Get("pn")
	pn, err := strconv.ParseInt(pnStr,10,64)
	if err!= nil || pn < 1 {
		pn = 1
	}
	hp.Data["Prev"], hp.Data["Next"], hp.Data["List"] =
		0,0,nil

	hp.Layout = "homeLayout.html"	//默认views下面，路径不用加views/
	hp.TplName = "home.html"
	err = hp.Render()
	fmt.Println(err)
}