package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/painterQ/painterBlog/models"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (lc *LoginController) URLMapping(){
	lc.Mapping("Get", lc.Get)
	lc.Mapping("Post", lc.Post)
}

// @router / [get]
func (lc *LoginController) Get() {
	lc.Data["Version"] = "1.0.1" //StaticVersion(c)
	lc.Data["BTitle"] = "login"

	isLogoutStr := lc.Input().Get("logout")
	isLogout, err := strconv.ParseBool(isLogoutStr)
	haveToken := checkToken(lc.Ctx)
	if (err != nil || isLogout ) && haveToken{
		lc.SetSecureCookie(string(macKey), "tk", "false")
		lc.Redirect("/admin/login",302)
		return
	}else if haveToken{
		lc.Redirect("/admin/profile",302)
		return
	}

	lc.TplName = "admin/login.html"
	err = lc.Render()
}

// @router / [post]
func (lc *LoginController) Post() {
	user := lc.Input().Get("user")
	pwd := lc.Input().Get("password")
	if user != pwd {
		//lc.Ctx.Output.Status = 400
		lc.SetSecureCookie(string(macKey), "tk", "",0, "/", nil, true)
		lc.Redirect("/admin/login",302)
		return
	}
	tk, ok := lc.GetSecureCookie(string(macKey), "tk")
	if ok && tk == models.ManagerSingleCase.Token() {
		lc.Redirect("/admin/profile", 302)
	}
	t := time.Now().Format(time.ANSIC)
	//age, domain, security, httpOnly
	lc.SetSecureCookie(string(macKey), "tk", "painterBlog"+t,
		/*-1, "/", nil, true*/)
	lc.Redirect("/admin/profile",302)
}

func CheckToken(ctx *context.Context) bool {
	if ctx.Request.URL.Path == "/admin/login"{
		return true
	}
	return checkToken(ctx)
}

func checkToken(ctx *context.Context) bool  {
	tk, ok := ctx.GetSecureCookie(string(macKey), "tk")
	if !ok {
		fmt.Println("@@@1")
		return false
	}
	if !strings.HasPrefix(tk, "painterBlog") || len(tk) < 12 {
		fmt.Println("@@@2")
		return false
	}
	if t, err := time.Parse(time.ANSIC,tk[11:]); err != nil {
		fmt.Println("@@@3",tk,err)
		return false
	} else {
		if time.Now().Sub(t) > time.Hour {
			fmt.Println("@@@4")
			return false
		}
	}
	err := models.ManagerSingleCase.NewUpdate().
		Update("lastTime",time.Now()).EndUpdate()
	if err != nil {
		logs.Error("更新登陆时间失败："+err.Error())
	}
	return true
}
