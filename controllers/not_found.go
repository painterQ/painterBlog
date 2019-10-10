package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type ErrorController struct {
	beego.Controller
}

func (e *ErrorController)Error404() {
	type Disqus struct { // 获取文章数量相关
		ShortName    string
		PublicKey    string
		AccessToken  string
		PostsCount   string
		PostsList    string
		PostCreate   string
		PostApprove  string
		ThreadCreate string
		Embed        string
		Interval     int
	}
	type Qiniu struct { // 七牛CDN
		Bucket    string
		Domain    string
		AccessKey string
		SecretKey string
	}
	type Twitter struct { // twitter信息
		Card    string
		Site    string
		Image   string
		Address string
	}
	e.Data = map[interface{}]interface{}{
		"BlogName": "BlogName",
		"SubTitle": "Ei.SubTitle",
		"Twitter":  new(Twitter),
		"CopyYear": time.Now().Year(),
		"BTitle":   "Ei.BTitle",
		"BeiAn":    "Ei.BeiAn",
		"Domain":   "localhost:8080",
		"Qiniu":    new(Qiniu),
		"Disqus":   new(Disqus),
	}

	e.Data["Version"] = "1.0.1"
	e.Data["Title"] = "Not Found"
	e.Data["Description"] = "404 Not Found"
	e.Data["Path"] = ""

	e.Layout = "homeLayout.html" //默认views下面，路径不用加views/
	e.TplName = "404.html"
	err := e.Render()
	fmt.Println(err)
}