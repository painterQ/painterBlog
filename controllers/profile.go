package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"strconv"
	"time"
)

type ProfileController struct {
	beego.Controller
}

func (pc *ProfileController) URLMapping(){
	pc.Mapping("Profile", pc.Profile)
	pc.Mapping("WritePost", pc.WritePost)
	pc.Mapping("ManagePosts",pc.ManagePosts)
}

// @router / [get]
func (pc *ProfileController) Profile() {
	pc.Data["Author"] = "PainterAuthor"
	pc.Data["Qiniu"] = models.CDNSingleCase
	pc.Data["Console"] = true
	pc.Data["Path"] = pc.Ctx.Request.URL.Path   // /admin/profile
	pc.Data["Title"] = "个人配置 | " + models.BlogSingleCase.BTitle()
	pc.Data["Manager"] = models.ManagerSingleCase
	pc.Data["Version"] = "1.0.1" //StaticVersion(c)
	pc.Data["Blog"] = models.BlogSingleCase

	isLogoutStr := pc.Input().Get("logout")
	isLogout, err := strconv.ParseBool(isLogoutStr)
	if err == nil || isLogout {
		//todo logout
		return
	}
	pc.Layout = "admin/backLayout.html"
	pc.TplName = "admin/profile.html"

	err = pc.Render()
	if err != nil{
		logs.Error("render err:",err.Error())
	}
}

// @router /write-post [get]
func (pc *ProfileController) WritePost() {
	pc.Data["Author"] = "PainterAuthor"
	pc.Data["Qiniu"] = models.CDNSingleCase

	idStr := pc.Input().Get("cid")
	id,err := strconv.ParseInt(idStr,10,64)
	if err == nil && id > 0 {
		artc := &models.Article{ID:id}
		qerr := models.Query(artc)
		if qerr !=nil {
			pc.Data["Title"] = "编辑文章 | " + models.BlogSingleCase.BTitle()
			pc.Data["Edit"] = artc
		}
	}

	pc.Data["Series"] = "Ei.Series"
	pc.Data["Path"] = pc.Ctx.Request.URL.Path   // /admin/profile
	pc.Data["Domain"] = "localhost:8080"

	//var tags []T
	//for tag := range Ei.Tags {
	//	tags = append(tags, T{tag, tag})
	//}
	//str, _ := json.Marshal(tags)
	pc.Data["Tags"] = string("tag1, tag2, tag3")

	pc.Layout = "admin/backLayout.html"
	pc.TplName = "admin/post.html"

	err = pc.Render()
	if err != nil{
		logs.Error("render err:",err.Error())
	}
}


//ManagePosts 文章管理==>Manage
// @router /manage-posts [get]
func (pc *ProfileController)ManagePosts() {
	kw := pc.Input().Get("keywords")
	serie,err := strconv.Atoi(pc.Input().Get("serie"))
	if err != nil || serie < 1 {
		serie = 0
	}
	page, err := strconv.Atoi(pc.Input().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	vals := pc.Ctx.Request.URL.Query()
	pc.Data["Author"] = models.ManagerSingleCase.Username()
	pc.Data["Qiniu"] = models.CDNSingleCase
	pc.Data["Manage"] = true  //是Manage不是Manager
	pc.Data["Path"] = pc.Ctx.Request.URL.Path
	pc.Data["Title"] = "文章管理 | " + models.BlogSingleCase.BTitle()
	//todo Series
	//pc.Data["Series"] = Ei.Series
	pc.Data["Serie"] = serie
	pc.Data["KW"] = kw
	var max int
	list, err := models.QueryArt(0,int(time.Now().Unix()))
	pc.Data["List"],max = list,len(list)
	if err != nil{
		logs.Error("加载文章错",err)
		return
	}

	if page < max {
		vals.Set("page", fmt.Sprint(page+1))
		pc.Data["Next"] = vals.Encode()
	}
	if page > 1 {
		vals.Set("page", fmt.Sprint(page-1))
		pc.Data["Prev"] = vals.Encode()
	}
	pc.Data["PP"] = make(map[int]string, max)
	for i := 0; i < max; i++ {
		vals.Set("page", fmt.Sprint(i+1))
		pc.Data["PP"].(map[int]string)[i+1] = vals.Encode()
	}
	pc.Data["Cur"] = page
	
	pc.Layout = "admin/backLayout.html"
	pc.TplName = "admin/posts.html"

	err = pc.Render()
	if err != nil{
		logs.Error("render err:",err.Error())
	}
}

