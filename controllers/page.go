package controllers

import (
	"github.com/astaxie/beego"
	"github.com/painterQ/painterBlog/models"
)


type PageController struct {
	beego.Controller
}

//URLMapping /
func (pc *PageController) URLMapping()  {
	pc.Mapping("GetIndex",pc.GetIndex)  // get /
	pc.Mapping("GetBackground",pc.GetBackground)  // get /background
}

// @router / [get]
func (pc *PageController) GetIndex()  {
	FriendLinks := models.AuthorSingleCase.Config.GetLinks()
	pc.Data["FriendLinks"] = FriendLinks
	pc.TplName = "index.html"
	err := pc.Render()
	if err != nil{
		responseJson(pc.Ctx,err)
		return
	}
}

// @router /background [get]
func (pc *PageController) GetBackground()  {
	pc.TplName = "background.html"
	err := pc.Render()
	if err != nil{
		responseJson(pc.Ctx,err)
		return
	}
}