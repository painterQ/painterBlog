package controllers

import (
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/utils"
)

var macKey = []byte{'0','2','5','5','1','9'}

func init() {
	beego.SetStaticPath("static","static")
	beego.AddTemplateExt("js")
	beego.AddTemplateExt("css")
	_ = beego.AddFuncMap("dateformat",utils.DateFormat)
	_ = beego.AddFuncMap("str2html",utils.Str2html)
	_ = beego.AddFuncMap("join",utils.Join)
	_ = beego.AddFuncMap("isnotzero",utils.IsNotZero)
	_ = beego.AddFuncMap("getavatar",utils.GetAvatar)

	//err := beego.AddViewPath("./views/admin")
	//if err !=nil{
	//	panic(err)
	//}
}